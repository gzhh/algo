package sort

func HeapSort(nums []int)  {
	n := len(nums)
	// 1.从完全二叉树的最后一个叶子结点的父结点开始遍历，将二叉树调整为一个最大堆
	for i := n / 2 - 1; i >= 0; i-- {
		maxHeapify(nums, i, n-1)
	}

	// 2.重复从最大堆取出数值最大的结点(把根结点和最后一个结点交换，把交换后的最后一个结点移出堆)，并让残余的堆维持最大堆性质
	for i := n - 1; i > 0; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		maxHeapify(nums, 0, i - 1)
	}
}

func maxHeapify(nums []int, start, end int)  {
	dad := start
	son := start * 2 + 1
	if son > end {
		return
	}

	if son + 1 < end && nums[son] <= nums[son+1] {
		son++
	}

	if nums[son] > nums[dad] {
		nums[son], nums[dad] = nums[dad], nums[son]
		maxHeapify(nums, son, end)
	}
}