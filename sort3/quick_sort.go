package sort3

func QuickSort(nums []int) []int {
	return decompose(nums, 0, len(nums)-1)
}

func decompose(nums []int, p int, q int) []int {
	if p >= q {
		return nums
	}
	// 需要注意，分区点不参与下次的分区，因为分区是已经确定比该值大或者小
	pivot := partition(nums, p, q)
	decompose(nums, p, pivot-1)
	decompose(nums, pivot+1, q)
	return nums
}

func partition(nums []int, p int, q int) int {
	pivot := nums[q]
	i, j := p, p
	for j < q {
		if nums[j] < pivot {
			Swap(nums, i, j)
			i++
		}
		j++
	}
	// swap for pivot
	Swap(nums, i, q)
	return i
}
