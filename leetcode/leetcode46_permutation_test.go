package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// 46. 全排列 https://leetcode.cn/problems/permutations/description/

// 非去重版本
func permute(nums []int) [][]int {
	indices := make([]int, 0)
	for index := range nums {
		indices = append(indices, index)
	}
	result := make([][]int, 0)
	result = permutation(nums, indices, result, make([]int, 0))
	return result
}

func permutation(nums []int, indices []int, result [][]int, selectedIndex []int) [][]int {
	if len(selectedIndex) == len(nums) {
		selected := make([]int, 0)
		for _, index := range selectedIndex {
			selected = append(selected, nums[index])
		}
		result = append(result, selected)
	}
	unselectedIndex := difference(indices, selectedIndex)
	for _, index := range unselectedIndex {
		// 去重
		copyTo := append([]int{}, selectedIndex...)
		copyTo = append(copyTo, index)
		result = permutation(nums, indices, result, copyTo)
	}
	return result
}

// difference 差集
func difference(o []int, t []int) []int {
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

func TestPermute(t *testing.T) {
	n1 := []int{1, 2, 3}
	r1 := permute(n1)
	assert.Equal(t,
		[][]int{
			{1, 2, 3},
			{1, 3, 2},
			{2, 1, 3},
			{2, 3, 1},
			{3, 1, 2},
			{3, 2, 1}},
		r1,
	)

	n2 := []int{0, 1}
	r2 := permute(n2)
	assert.Equal(t,
		[][]int{
			{0, 1},
			{1, 0}},
		r2,
	)

	n3 := []int{1}
	r3 := permute(n3)
	assert.Equal(t, [][]int{{1}}, r3)
}
