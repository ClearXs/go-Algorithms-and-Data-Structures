package sort3

func MergeSort(nums []int) []int {
	return resolve(nums, 0, len(nums)-1)
}

func resolve(nums []int, p int, q int) []int {
	// 递归终止条件
	if p >= q {
		return nums[p : q+1]
	}
	// 取中间点
	mid := (p + q) / 2
	// left
	left := resolve(nums, p, mid)
	// right
	right := resolve(nums, mid+1, q)
	return merge(left, right)
}

// 合并数组
// left 左数组
// right 右数组
func merge(left []int, right []int) []int {
	capacity := len(left) + len(right)
	// 利用tmp数组，把left 与 right按照其大小进行合并
	tmp := make([]int, capacity)
	i, j, k := 0, 0, 0
	for i < len(left) || j < len(right) {
		if i == len(left) {
			tmp[k] = right[j]
			j++
		} else if j == len(right) {
			tmp[k] = left[i]
			i++
		} else if left[i] < right[j] {
			tmp[k] = left[i]
			i++
		} else {
			tmp[k] = right[j]
			j++
		}
		k++
	}
	return tmp
}
