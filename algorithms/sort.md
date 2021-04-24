# Sort Algorithms
### Bubble Sort [wiki](https://en.wikipedia.org/wiki/Bubble_sort)

**Golang Solution**
```go
func bubbleSort(nums []int)  {
	n := len(nums)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
}
```