package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// 47. 全排列 II https://leetcode.cn/problems/permutations-ii/description/
// 给定一个可包含重复数字的序列 nums ，按任意顺序 返回所有不重复的全排列。

// 去重版本
func permuteUnique(nums []int) [][]int {
	indices := make([]int, 0)
	for index := range nums {
		indices = append(indices, index)
	}
	result := make([][]int, 0)
	result = permutationForUnique(nums, indices, result, make([]int, 0))
	return result
}

func permutationForUnique(nums []int, indices []int, result [][]int, selectedIndex []int) [][]int {
	if len(selectedIndex) == len(nums) {
		selected := make([]int, 0)
		for _, index := range selectedIndex {
			selected = append(selected, nums[index])
		}
		result = append(result, selected)
	}
	unselectedIndex := differenceUnique(indices, selectedIndex)
	duplicates := make(map[int]int)
	for _, index := range unselectedIndex {
		// 去重
		num := nums[index]
		if _, ok := duplicates[num]; !ok {
			copyTo := append([]int{}, selectedIndex...)
			copyTo = append(copyTo, index)
			result = permutationForUnique(nums, indices, result, copyTo)
			duplicates[num] = num
		}
	}
	return result
}

// differenceUnique 差集
func differenceUnique(o []int, t []int) []int {
	differ := make([]int, 0)
	tMap := make(map[int]any)
	for _, v := range t {
		tMap[v] = v
	}
	for _, v := range o {
		if _, ok := tMap[v]; !ok {
			differ = append(differ, v)
		}
	}
	return differ
}

func TestPermuteUnique(t *testing.T) {
	n1 := []int{1, 1, 2}
	r1 := permuteUnique(n1)
	assert.Equal(t,
		[][]int{
			{1, 1, 2},
			{1, 2, 1},
			{2, 1, 1}}, r1,
	)

	n2 := []int{1, 2, 3}
	r2 := permuteUnique(n2)
	assert.Equal(t,
		[][]int{{1, 2, 3},
			{1, 3, 2},
			{2, 1, 3},
			{2, 3, 1},
			{3, 1, 2},
			{3, 2, 1}}, r2,
	)
}
