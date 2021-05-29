package slidingwindow

import (
	"testing"
)

func TestMinSubArrayLen(t *testing.T) {
	var numsArray = [][]int{
		{2, 3, 1, 2, 4, 3},
		{1, 4, 4},
		{1, 1, 1, 1, 1, 1, 1, 1},
	}
	var targets = []int{7, 4, 11}
	var wants = []int{2, 1, 0}
	for i, want := range wants {
		if got := MinSubArrayLen(targets[i], numsArray[i]); got != want {
			t.Fatalf("MinSubArrayLen(nums, %v) = %v, want %v", numsArray[i], got, want)
		}
	}
}
