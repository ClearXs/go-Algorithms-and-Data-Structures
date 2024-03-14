// https://leetcode.cn/problems/two-sum/
// 两数之和
// 给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。
//
// 你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。
//
// 你可以按任意顺序返回答案。
package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSumOfTow(t *testing.T) {
	// 采取hash表数据数据，以key作为数组索引值，value作为索引
	r1 := towSums([]int{2, 7, 11, 15}, 9)
	assert.Equal(t, r1, []int{0, 1})

	r2 := towSums([]int{3, 3, 3}, 6)
	assert.Equal(t, r2, []int{0, 1})
}

func towSums(nums []int, target int) []int {
	hash := make(map[int]int, len(nums))
	for i := range nums {
		s := target - nums[i]
		// 程序顺序写法问题，数据查找后置，只会查找出先匹配的数据
		if value, ok := hash[s]; ok {
			return []int{value, i}
		}
		hash[nums[i]] = i
	}
	return nil
}
