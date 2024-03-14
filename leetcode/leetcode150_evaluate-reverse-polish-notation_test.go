package leetcode

import (
	"data-structure/stack"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

// 150. 逆波兰表达式求值 https://leetcode.cn/problems/evaluate-reverse-polish-notation/description/
// 给你一个字符串数组 tokens ，表示一个根据 逆波兰表示法 表示的算术表达式。
//
// 请你计算该表达式。返回一个表示表达式值的整数。

const (
	additionToken    = "+"
	subtractionToken = "-"
	multiplyToken    = "*"
	divideToken      = "/"
)

func evalRPN(tokens []string) int {
	s := stack.New[string]()

	eval := func(s *stack.Stack[string]) (int, int) {
		numeric1, err := strconv.Atoi(*s.Pop())
		if err != nil {
			panic(err)
		}
		numeric2, err := strconv.Atoi(*s.Pop())
		if err != nil {
			panic(err)
		}
		return numeric1, numeric2
	}
	for _, token := range tokens {
		if token == additionToken {
			numeric1, numeric2 := eval(s)
			s.Push(strconv.Itoa(numeric1 + numeric2))
		} else if token == subtractionToken {
			numeric1, numeric2 := eval(s)
			s.Push(strconv.Itoa(numeric2 - numeric1))
		} else if token == multiplyToken {
			numeric1, numeric2 := eval(s)
			s.Push(strconv.Itoa(numeric1 * numeric2))
		} else if token == divideToken {
			numeric1, numeric2 := eval(s)
			s.Push(strconv.Itoa(numeric2 / numeric1))
		} else {
			s.Push(token)
		}
	}
	result, err := strconv.Atoi(*s.Pop())
	if err != nil {
		panic(err)
	}
	return result
}

func TestEvalRPN(t *testing.T) {
	tokens1 := []string{"2", "1", "+", "3", "*"}
	r1 := evalRPN(tokens1)
	assert.Equal(t, 9, r1)

	tokens2 := []string{"4", "13", "5", "/", "+"}
	r2 := evalRPN(tokens2)
	assert.Equal(t, 6, r2)

	tokens3 := []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}
	r3 := evalRPN(tokens3)
	assert.Equal(t, 22, r3)
}
