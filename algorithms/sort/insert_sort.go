package sort

func InsertSort(nums []int) {
	n := len(nums)
	for i := 1; i < n; i++ {
		num := nums[i]
		j := i - 1
		for j >= 0 && nums[j] > num {
			nums[j+1] = nums[j]
			j--
		}
		nums[j+1] = num
	}
}
