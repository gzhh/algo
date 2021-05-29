package dynamicprogramming

import (
	"testing"
)

func TestClimbStairsV2(t *testing.T) {
	var nums = []int{1, 2, 3, 4, 5}
	var wants = []int{1, 2, 4, 7, 13}
	for i, want := range wants {
		if got := ClimbStairsV2(nums[i]); got != want {
			t.Fatalf("ClimbStairsV2(num, %v) = %v, want %v", nums[i], got, want)
		}
	}
}
