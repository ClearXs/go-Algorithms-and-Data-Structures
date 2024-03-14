package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// 131. 分割回文串 https://leetcode.cn/problems/palindrome-partitioning/description/

// 给你一个字符串 s，请你将 s 分割成一些子串，使每个子串都是 回文串 。返回 s 所有可能的分割方案。
//
//	回文串 是正着读和反着读都一样的字符串。

func partitionPalindrome(s string) [][]string {
	result := make([][]string, 0)
	result = spliterator(s, result, make([]string, 0), 0)
	return result
}

func spliterator(s string, result [][]string, splits []string, startIndex int) [][]string {
	// 递归结束条件
	if startIndex > len(s)-1 {
		if len(splits) > 0 {
			// 合并数据
			result = append(result, append([]string{}, splits...))
		}
		return result
	}
	for i := startIndex; i <= len(s)-1; i++ {
		newSplits := append([]string{}, splits...)
		// 在一次回溯中，对每一层的分割后的字串满足回文则放到分割数组中
		// 分割
		cuttingBytes := []byte(s)[startIndex : i+1]
		cutStr := string(cuttingBytes)
		if isPalindrome(cutStr) {
			newSplits = append(newSplits, cutStr)
		} else {
			// 非回文串那么本次分割作废
			continue
		}
		result = spliterator(s, result, newSplits, i+1)
	}
	return result
}

// 判断是否回文
func isPalindrome(s string) bool {
	bytes := []byte(s)
	startIndex, endIndex := 0, len(bytes)-1
	for endIndex >= startIndex {
		if bytes[endIndex] != bytes[startIndex] {
			return false
		}
		startIndex++
		endIndex--
	}
	return true
}

func TestPartitionPalindrome(t *testing.T) {
	s1 := "aab"
	r1 := partitionPalindrome(s1)
	assert.Equal(t, [][]string{{"a", "a", "b"}, {"aa", "b"}}, r1)

	s2 := "a"
	r2 := partitionPalindrome(s2)
	assert.Equal(t, [][]string{{"a"}}, r2)

	s3 := "cdd"
	r3 := partitionPalindrome(s3)
	assert.Equal(t, [][]string{{"c", "d", "d"}, {"c", "dd"}}, r3)
}
