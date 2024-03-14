package leetcode

import (
	"data-structure/stack"
	"github.com/stretchr/testify/assert"
	"testing"
)

// 1047. 删除字符串中的所有相邻重复项 https://leetcode.cn/problems/remove-all-adjacent-duplicates-in-string/description/
// 给出由小写字母组成的字符串 S，重复项删除操作会选择两个相邻且相同的字母，并删除它们。
//
// 在 S 上反复执行重复项删除操作，直到无法继续删除。
//
// 在完成所有重复项删除操作后返回最终的字符串。答案保证唯一。

func removeStringsDuplicates(s string) string {
	stk := stack.New[byte]()
	bytes := []byte(s)
	for _, char := range bytes {
		last := stk.PeekLast()
		if last != nil && *last == char {
			stk.Pop()
		} else {
			stk.Push(char)
		}
	}
	return string(stk.GetItems())
}

func TestRemoveStringsDuplicates(t *testing.T) {
	s := "abbaca"
	d1 := removeStringsDuplicates(s)
	assert.Equal(t, "ca", d1)
}
