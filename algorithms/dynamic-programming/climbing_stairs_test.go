package dynamicprogramming

import (
	"testing"
)

func TestClimbStairs(t *testing.T) {
	var nums = []int{1, 2, 3, 4, 5}
	var wants = []int{1, 2, 3, 5, 8}
	for i, want := range wants {
		if got := ClimbStairs(nums[i]); got != want {
			t.Fatalf("ClimbStairs(num, %v) = %v, want %v", nums[i], got, want)
		}
	}
}
