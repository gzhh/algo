package dynamicprogramming

// lcci08.01: https://leetcode-cn.com/problems/three-steps-problem-lcci/
func ClimbStairsV2(n int) int {
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		if i == 1 {
			dp[i] = 1
		} else if i == 2 {
			dp[i] = 2
		} else if i == 3 {
			dp[i] = 4
		} else {
			tmp := (dp[i-3] + dp[i-2]) % 1000000007
			dp[i] = (tmp + dp[i-1]) % 1000000007
		}
	}
	return dp[n]
}
