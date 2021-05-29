package slidingwindow

import "math"

// Solution1: sliding window
func MinSubArrayLen(s int, nums []int) int {
	ans := math.MaxInt32
	sum := 0
	start, end := 0, 0
	for end < len(nums) {
		sum += nums[end]
		for sum >= s {
			ans = min(ans, end-start+1)
			sum -= nums[start]
			start++
		}
		end++
	}

	if ans == math.MaxInt32 {
		return 0
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Solution2: 前缀和 + 二分查找
// TODO figure out
