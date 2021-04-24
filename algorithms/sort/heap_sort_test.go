package sort

import (
	"testing"
)

func TestHeapSort(t *testing.T)  {
	var nums = []int{4, 1, 8, 5, 3, 9, 2, 7, 3, 6}
	var wantNums = []int{1, 2, 3, 3, 4, 5, 6, 7, 8, 9}
	HeapSort(nums)
	for key, num := range nums {
		if num != wantNums[key] {
			t.Fatalf("HeapSort(nums) = %v, want %v", nums, wantNums)
		}
	}
}