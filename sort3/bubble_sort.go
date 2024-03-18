package sort3

func BubbleSort(nums []int) []int {
	capacity := len(nums)
	if capacity <= 0 {
		return nums
	}
	for i, _ := range nums {
		j := 0
		// 内存循环终止条件为什么是capacity-i-1
		// 1.每次冒泡的结果使得某个元素在数组最后一个 所以-i
		// 2.-1目的保证数组不越位
		for j < capacity-i-1 {
			if nums[j] > nums[j+1] {
				swap := nums[j+1]
				nums[j+1] = nums[j]
				nums[j] = swap
			}
			j++
		}
	}
	return nums
}
