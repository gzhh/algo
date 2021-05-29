package dynamicprogramming

// ref: https://leetcode-cn.com/problems/maximum-sum-circular-subarray/solution/huan-xing-zi-shu-zu-de-zui-da-he-by-leetcode/
// TODO figure out
func MaxSubarraySum(nums []int) int {
	n := len(nums)

	last := nums[0]
	maxVal := nums[0]
	sum := nums[0]
	for i := 1; i < n; i++ {
		last = nums[i] + max(last, 0)
		maxVal = max(maxVal, last)
		sum += nums[i]
	}

	last = 0
	minVal := 0
	for i := 1; i < n-1; i++ {
		last = nums[i] + min(last, 0)
		minVal = min(minVal, last)
	}

	return max(maxVal, sum-minVal)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
