// 27. 移除元素
// https://leetcode.cn/problems/remove-element/
// 给你一个数组 nums 和一个值 val，你需要 原地 移除所有数值等于 val 的元素，并返回移除后数组的新长度。
//
// 不要使用额外的数组空间，你必须仅使用 O(1) 额外空间并 原地 修改输入数组。
//
// 元素的顺序可以改变。你不需要考虑数组中超出新长度后面的元素。
//
// 说明:
//
// 为什么返回数值是整数，但输出的答案是数组呢?
//
// 请注意，输入数组是以「引用」方式传递的，这意味着在函数里修改输入数组对于调用者是可见的。
//
// 你可以想象内部操作如下:
//
// // nums 是以“引用”方式传递的。也就是说，不对实参作任何拷贝
// int len = removeElement(nums, val);
//
// // 在函数里修改输入数组对于调用者是可见的。
// // 根据你的函数返回的长度, 它会打印出数组中 该长度范围内 的所有元素。
//
//	for (int i = 0; i < len; i++) {
//	   print(nums[i]);
//	}
package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	assert.Equal(t, 2, removeElement([]int{3, 2, 2, 3}, 3))
	assert.Equal(t, 5, removeElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2))

}

// 根据题目要求，必须要是原地算法，所以需要考虑多个指针，根据题目理解，需要返回的数组必须不包含重复元素，这就可以采用快慢双指针
func removeElement(nums []int, val int) int {
	slow, fast := 0, 0
	for fast != len(nums) {
		if nums[fast] != val {
			swap := nums[fast]
			nums[fast] = nums[slow]
			nums[slow] = swap
			slow++
		}
		fast++
	}
	return slow
}
