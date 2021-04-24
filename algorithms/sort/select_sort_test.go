package sort

import (
	"testing"
)

func TestSelectSort(t *testing.T)  {
	var nums = []int{4, 1, 8, 5, 3, 9, 2, 7, 3, 6}
	var wantNums = []int{1, 2, 3, 3, 4, 5, 6, 7, 8, 9}
	SelectSort(nums)
	for key, num := range nums {
		if num != wantNums[key] {
			t.Fatalf("SelectSort(nums) = %v, want %v", nums, wantNums)
		}
	}
}