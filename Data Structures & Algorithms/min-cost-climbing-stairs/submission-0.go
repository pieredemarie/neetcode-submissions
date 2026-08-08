func minCostClimbingStairs(cost []int) int {
    dp := make([]int, len(cost)+1)

	dp[0] = cost[0]
	dp[1] = cost[1]

	n := len(cost)

	for i := 2;i<n;i++ {
		dp[i] = cost[i] + min(dp[i-2],dp[i-1])
	}

	return min(dp[n-1],dp[n-2])
}
