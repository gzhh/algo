package search

import (
	"testing"
)

func TestBinarySearch(t *testing.T)  {
	var nums = []int{1, 2, 4, 5, 6, 7, 9}
	var targets = []int{1, 2, 4, 5, 6, 7, 9}
	var wants = []int{0, 1, 2, 3, 4, 5, 6}
	for i, want := range wants {
		if get := BinarySearch(nums, targets[i]); get != want {
			t.Fatalf("BinarySearch(nums, %v) = %v, want %v", targets[i], get, want)
		}
	}
}