package dynamicprogramming

import (
	"testing"
)

func TestMaxSubarraySumCircular(t *testing.T) {
	var numsArray = [][]int{
		{1, -2, 3, -2},
		{5, -3, 5},
		{3, -1, 2, -1},
		{3, -2, 2, -3},
		{-2, -3, -1},
	}
	var wants = []int{3, 10, 4, 3, -1}
	for i, want := range wants {
		if got := MaxSubarraySum(numsArray[i]); got != want {
			t.Fatalf("MaxSubarraySum(nums, %v) = %v, want %v", numsArray[i], got, want)
		}
	}
}
