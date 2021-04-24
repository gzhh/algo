package sort

func BubbleSort(nums []int)  {
	n := len(nums)
	for i := 0; i < n; i++ {
		for j := 0; j < n - i - 1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
}