package sort3

func SelectionSort(nums []int) []int {
	capacity := len(nums)
	if capacity <= 1 {
		return nums
	}

	for i, _ := range nums {
		minIndex := i
		innerIndex := i
		for innerIndex < capacity {
			if nums[innerIndex] < nums[minIndex] {
				minIndex = innerIndex
			}
			innerIndex++
		}
		// swap
		swap := nums[i]
		nums[i] = nums[minIndex]
		nums[minIndex] = swap
	}
	return nums
}
