package leetcode

import "sort"

// solution1
func smallestK(arr []int, k int) []int {
	sort.Ints(arr)
	return arr[:k]
}

// solution2
func smallestK(arr []int, k int) []int {
	if len(arr) <= k {
		return arr
	}
	if k == 0 {
		return []int{}
	}
	heap := arr[:k]
	for i := k / 2; i >= 1; i-- {
		heapify(heap, k, i)
	}
	for i := len(arr) - 1; i >= k; i-- {
		if arr[i] < heap[0] {
			heap[0] = arr[i]
			heapify(heap, k, 1)
		}
	}
	return heap
}

func heapify(arr []int, k int, i int) {
	for i < k {
		max := i
		n := i * 2
		if n <= k && arr[i-1] < arr[n-1] {
			max = n
		}
		if n+1 <= k && arr[max-1] < arr[n] {
			max = n + 1
		}
		if max == i {
			break
		}
		arr[i-1], arr[max-1] = arr[max-1], arr[i-1]
		i = max
	}
}
