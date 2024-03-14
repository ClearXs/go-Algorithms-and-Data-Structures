// 15. 三数之和
// https://leetcode.cn/problems/3sum/
// 给你一个整数数组 nums ，判断是否存在三元组 [nums[i], nums[j], nums[k]] 满足 i != j、i != k 且 j != k ，同时还满足 nums[i] + nums[j] + nums[k] == 0 。请
//
// 你返回所有和为 0 且不重复的三元组。
//
// 注意：答案中不可以包含重复的三元组。
package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestThreeSum(t *testing.T) {
	assert.Equal(t, 2, len(threeSum([]int{-1, 0, 1, 2, -1, -4}, 0)))
}

func threeSum(nums []int, target int) [][]int {
	sort(nums)
	return nSum(nums, 3, target, make([][]int, 0))
}

// nSum 求解n数之和等target
// nums 必须是已经排序的数组
func nSum(nums []int, n int, target int, result [][]int) [][]int {
	if len(nums) < n {
		return result
	}
	// 要求返回数据的下标不一致的，如[nums[a], nums[b], nums[c]] a != b != c
	// 证明：
	// 初始化：当n=2时，返回的数组下标不一致
	// 当n=3时，数组下标i 与递归n=2的数组下标不一致
	// 当n=k时，每次循环的数组下标与内层循环都不一致
	// 所以当n=k+1时，得到结果的数组下标不一致

	// 多数组和：当数组内部存在多个数据与目标数组和一致，则返回多个数据数组和
	if n == 2 {
		// 2数和解法，采用左右指针法
		left, right := 0, len(nums)-1
		for left < right {
			sum := nums[left] + nums[right]
			if sum == target {
				result = append(result, []int{nums[left], nums[right]})
				left++
			} else if sum < target {
				left++
			} else if sum > target {
				right--
			}
		}
	} else {
		for index, num := range nums {
			// 重要
			// 下层递归中：
			// nums[index+1:]：数组需要往后移动一位
			// n-1: 层数需要-1
			// target-num: 目标数据需要是当前数据相减，如num[i] + nums[j] = target - nums[k]
			// make([][]int, 0): 构造多数据和
			r := nSum(nums[index+1:], n-1, target-num, make([][]int, 0))
			// 判断是否存在相等
			if len(r) > 0 {
				for rt := range r {
					r[rt] = append(r[rt], num)
					// 去重并添加到返回的结果集中
					// hash去重
					if len(result) == 0 {
						result = append(result, r[rt])
					} else {
						ok := false
						for _, ele := range result {
							dp := duplicate(r[rt], ele)
							ok = ok || dp
						}
						if !ok {
							result = append(result, r[rt])
						}
					}
				}
			}
		}
	}
	return result
}

// 插入排序
func sort(nums []int) {
	for i := range nums {
		j := i
		// 采用尾插法
		for j > 0 {
			if nums[j] < nums[j-1] {
				swap := nums[j]
				nums[j] = nums[j-1]
				nums[j-1] = swap
			}
			j--
		}
	}
}

// duplicate true 是重复 false 不重复
func duplicate(n1 []int, n2 []int) bool {
	dup := make(map[int]int)
	for index, num := range n1 {
		dup[num] = index
	}
	for _, ele := range n2 {
		if _, ok := dup[ele]; !ok {
			return false
		}
	}
	return true
}
