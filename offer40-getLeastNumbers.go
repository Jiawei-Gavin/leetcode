package leetcode

func getLeastNumbers(arr []int, k int) []int {
	if len(arr) == 0 || k == 0 {
		return nil
	}
	quicksort(arr, k, 0, len(arr)-1)
}

func quicksort(arr []int, k, left, right int) []int {
	i, j := left, right
	for i < j {
		for i < j && arr[j] >= arr[left] {
			j--
		}
		for i < j && arr[j] <= arr[left] {
			i++
		}
		arr[i], arr[j] = arr[j], arr[i]
	}
	arr[i], arr[left] = arr[left], arr[i]
	if i > k {
		return quicksort(arr, k, left, i-1)
	}
	if i < k {
		return quicksort(arr, k, i+1, right)
	}
	return arr[:k]
}
