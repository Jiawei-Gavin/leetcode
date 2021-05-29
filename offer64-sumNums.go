package leetcode

var result int

func sumNums(n int) int {
	result = 0
	sum(n)
	return result
}

func sum(n int) bool {
	result += n
	return 0 < n && sum(n-1)
}
