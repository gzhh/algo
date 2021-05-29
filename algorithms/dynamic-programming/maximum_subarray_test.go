package dynamicprogramming

import (
	"testing"
)

func TestMaxSubArray(t *testing.T) {
	var numsArray = [][]int{
		{-2, 1, -3, 4, -1, 2, 1, -5, 4},
		{1},
		{-1},
		{5, 4, -1, 7, 8},
	}
	var wants = []int{6, 1, -1, 23}
	for i, want := range wants {
		if got := MaxSubArray(numsArray[i]); got != want {
			t.Fatalf("MaxSubArray(nums, %v) = %v, want %v", numsArray[i], got, want)
		}
	}
}
