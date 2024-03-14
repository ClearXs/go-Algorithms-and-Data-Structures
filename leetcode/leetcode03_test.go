// https://leetcode.cn/problems/longest-substring-without-repeating-characters/
// 无重复字符的最长子串
// 给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。
package leetcode

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	assert.Equal(t, 3, lengthOfLongestSubstring("abcabcbb"))
	assert.Equal(t, 1, lengthOfLongestSubstring("bbbbb"))
	assert.Equal(t, 3, lengthOfLongestSubstring("pwwkew"))
	assert.Equal(t, 3, lengthOfLongestSubstring("abcacb"))
	assert.Equal(t, 1, lengthOfLongestSubstring(" "))
	assert.Equal(t, 2, lengthOfLongestSubstring("aab"))
}

func lengthOfLongestSubstring(s string) int {
	// 采用滑动窗口解题
	left, right, maxLength := 0, 0, 0
	bytes := []byte(s)
	str := []rune(string(bytes))
	for ; right < len(str); right++ {
		if right > 0 {
			c := str[right]
			substr := s[left:right]
			// 判断左侧窗口是否要收缩
			hasIndex := strings.IndexAny(substr, string(c))
			// 遇到重复字符，判断是否已经是最大子子字符串
			if hasIndex > -1 {
				thisLength := right - left
				if thisLength > maxLength {
					maxLength = thisLength
				}
				left = left + hasIndex + 1
			}
		}
	}

	if right > left && right-left > maxLength {
		maxLength = right - left
	}
	return maxLength
}
