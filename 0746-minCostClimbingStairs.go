package leetcode

func minCostClimbingStairs(cost []int) int {
	preCost, curCost := 0, 0
	for i := 2; i <= len(cost); i++ {
		preCost, curCost = curCost, min(curCost+cost[i-1], preCost+cost[i-2])
	}
	return curCost
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
