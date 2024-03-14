// 704. 二分查找
// https://leetcode.cn/problems/binary-search/
// 给定一个 n 个元素有序的（升序）整型数组 nums 和一个目标值 target  ，写一个函数搜索 nums 中的 target，如果目标值存在返回下标，否则返回 -1。
package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSearch(t *testing.T) {
	assert.Equal(t, 4, search([]int{-1, 0, 3, 5, 9, 12}, 9))
	assert.Equal(t, -1, search([]int{-1, 0, 3, 5, 9, 12}, 2))

}

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	// 左闭右闭区间
	for left <= right {
		// 计算中间值，这会避免整数溢出的可能（传统的(left + right) / 2将会导致）
		mid := left + (right-left)/2
		// 下面的 + 1 - 1在文章https://labuladong.github.io/algo/di-ling-zh-bfe1b/wo-xie-le--3c789/写得很清楚
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			// 往右区间找
			left = mid + 1
		} else if nums[mid] > target {
			// 往左区间找
			right = mid - 1
		}
	}
	return -1
}
