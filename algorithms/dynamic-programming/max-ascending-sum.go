package dynamicprogramming

// ref leetcode 53. Maximum Subarray
func MaxAscendingSum(nums []int) int {
	n := len(nums)
	dp := make([]int, n)

	dp[0] = nums[0]
	for i := 1; i < n; i++ {
		if dp[i-1]+nums[i] > nums[i] && nums[i-1] < nums[i] {
			dp[i] = dp[i-1] + nums[i]
		} else {
			dp[i] = nums[i]
		}
	}

	max := dp[0]
	for _, num := range dp {
		if num > max {
			max = num
		}
	}
	return max
}
