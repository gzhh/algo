package dynamicprogramming

import (
	"testing"
)

func TestMaxAscendingSum(t *testing.T) {
	var numsArray = [][]int{
		{10, 20, 30, 5, 10, 50},
		{10, 20, 30, 40, 50},
		{12, 17, 15, 13, 10, 11, 12},
		{100, 10, 1},
	}
	var wants = []int{65, 150, 33, 100}
	for i, want := range wants {
		if got := MaxAscendingSum(numsArray[i]); got != want {
			t.Fatalf("MaxAscendingSum(nums, %v) = %v, want %v", numsArray[i], got, want)
		}
	}
}
