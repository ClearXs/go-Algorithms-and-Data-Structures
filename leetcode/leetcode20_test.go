// https://leetcode.cn/problems/valid-parentheses/
// 有效的括号
// 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
//
// 有效字符串需满足：
//
// 左括号必须用相同类型的右括号闭合。
// 左括号必须以正确的顺序闭合。
// 每个右括号都有一个对应的相同类型的左括号
package leetcode

import (
	"data-structure/stack"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsValid(t *testing.T) {
	assert.Equal(t, true, isValid("()"))
	assert.Equal(t, true, isValid("()[]{}"))
	assert.Equal(t, false, isValid("(}"))
	assert.Equal(t, true, isValid("{[]}"))
}

// 栈数据结构

func isValid(s string) bool {
	bra := []rune(s)
	if len(bra) < 2 {
		return false
	}
	var bracket = make(map[string]string)
	// 初始化
	bracket["("] = ")"
	bracket["{"] = "}"
	bracket["["] = "]"
	// 采用栈数据结构
	// 经观察得知，左括号与右括号一定是"相邻"出现，如()、{()}，这样的数据结构采用栈没有问题
	s1 := stack.Stack[string]{}
	for _, val := range bra {
		v := string(val)
		if v == "(" || v == "[" || v == "{" {
			s1.Push(v)
		} else {
			left := s1.Pop()
			right := bracket[*left]
			if right != v {
				return false
			}
		}
	}
	return true
}
