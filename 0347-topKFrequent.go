package leetcode

import (
	"container/heap"
	"sort"
)

// solution1
func topKFrequent(nums []int, k int) []int {
	hashMap := map[int]int{}
	m := make([]int, 0)
	for _, num := range nums {
		i, ok := hashMap[num]
		if ok {
			hashMap[num] = i + 1
		} else {
			hashMap[num] = 1
			m = append(m, num)
		}
	}

	sort.Slice(m, func(i, j int) bool {
		return hashMap[m[i]] > hashMap[m[j]]
	})
	return m[:k]
}

// solution2
func topKFrequent(nums []int, k int) []int {
	hashMap := map[int]int{}
	for _, num := range nums {
		hashMap[num]++
	}

	h := &IHeap{}
	heap.Init(h)
	for key, value := range hashMap {
		heap.Push(h, [2]int{key, value})
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	result := make([]int, k)
	for i := 0; i < k; i++ {
		result[k-i-1] = heap.Pop(h).([2]int)[0]
	}
	return result
}

type IHeap [][2]int

func (h IHeap) Len() int           { return len(h) }
func (h IHeap) Less(i, j int) bool { return h[i][1] < h[j][1] }
func (h IHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IHeap) Push(x interface{}) {
	*h = append(*h, x.([2]int))
}

func (h *IHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
