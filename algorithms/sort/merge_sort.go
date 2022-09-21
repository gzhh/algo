package sort

func MergeSort(nums []int) {
	n := len(nums)
	tmpNums := make([]int, n)
	merge(nums, tmpNums, 0, n-1)
}

func merge(nums, tmpNums []int, start, end int) {
	if start >= end {
		return
	}

	// devide
	mid := start + (end-start)/2
	merge(nums, tmpNums, start, mid)
	merge(nums, tmpNums, mid+1, end)

	// conquer: merge two sorted list
	i := start
	j := mid + 1
	index := start
	for i <= mid && j <= end {
		if nums[i] < nums[j] {
			tmpNums[index] = nums[i]
			i++
		} else {
			tmpNums[index] = nums[j]
			j++
		}
		index++
	}
	for i <= mid {
		tmpNums[index] = nums[i]
		i++
		index++
	}
	for j <= end {
		tmpNums[index] = nums[j]
		j++
		index++
	}
	for i := start; i <= end; i++ {
		nums[i] = tmpNums[i]
	}
}
