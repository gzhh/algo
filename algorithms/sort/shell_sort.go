package sort

func ShellSort(nums []int) {
	n := len(nums)
	gap := n / 2
	for gap > 0 {
		for i := gap; i < n; i++ {
			num := nums[i]
			j := i - gap
			for j >= 0 && nums[j] > num {
				nums[j+gap] = nums[j]
				j -= gap
			}
			nums[j+gap] = num
		}
		gap /= 2
	}
}
