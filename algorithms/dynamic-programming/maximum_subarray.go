package dynamicprogramming

// dp[i] 代表以第 i 个数结尾的「连续子数组的最大和」，那么结果即为 max{dp[0], ..., dp[n-1]}
// 状态转移方程：dp[i] = max(dp[i-1] + nums[i], nums[i])
// Time O(n), Space O(n)
func MaxSubArray(nums []int) int {
	n := len(nums)
	dp := make([]int, n)

	dp[0] = nums[0]
	for i := 1; i < n; i++ {
		if dp[i-1]+nums[i] > nums[i] {
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

// Space O(n) => O(1)
func maxSubArray1(nums []int) int {
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i-1]+nums[i] > nums[i] {
			nums[i] += nums[i-1]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}

// another solution: devide and conquer
