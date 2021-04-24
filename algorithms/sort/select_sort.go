package sort

func SelectSort(nums []int)  {
	n := len(nums)
	for i := 0; i < n; i++ {
		key := i
		for j := i + 1; j < n; j++ {
			if nums[j] < nums[key] {
				key = j
			}
		}
		nums[i], nums[key] = nums[key], nums[i]
	}
}