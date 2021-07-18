package leetcode

func findDisappearedNumbers(nums []int) []int {
	n := len(nums)
	for _, num := range nums {
		temp := (num - 1) % n
		nums[temp] += n
	}
	var res []int
	for i := 0; i < n; i++ {
		if nums[i] <= n {
			res = append(res, i+1)
		}
	}
	return res
}
