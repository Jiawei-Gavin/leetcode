package leetcode

import "math"

func maxProfit(prices []int) int {
	cost, profit := math.MaxInt32, 0
	for _, price := range prices {
		cost = min(cost, price)
		profit = max(profit, price-cost)
	}
	return profit
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
