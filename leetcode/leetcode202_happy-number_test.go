package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// 202. 快乐数 https://leetcode.cn/problems/happy-number/description/

// 编写一个算法来判断一个数 n 是不是快乐数。
//
//「快乐数」 定义为：
//
// 对于一个正整数，每一次将该数替换为它每个位置上的数字的平方和。
// 然后重复这个过程直到这个数变为 1，也可能是 无限循环 但始终变不到 1。
// 如果这个过程 结果为 1，那么这个数就是快乐数。
// 如果 n 是 快乐数 就返回 true ；不是，则返回 false 。

func isHappy(n int) bool {
	duplicates := make(map[int]int)
	sum := n
	for {
		if _, ok := duplicates[sum]; ok {
			return false
		} else {
			duplicates[sum] = sum
		}
		sum = getSum(sum)
		if sum == 1 {
			return true
		}
	}
}

func getSum(num int) int {
	sum := 0
	nextNum := num
	for {
		if nextNum == 0 {
			return sum
		}
		// 求该数最后一位，如 82 % 10 = 2
		place := nextNum % 10
		sum += place * place
		// 求该数去除最后一位，如 82 / 10 = 8
		nextNum = nextNum / 10
	}
}

func TestIsHappy(t *testing.T) {
	r1 := isHappy(19)
	assert.Equal(t, true, r1)

	r2 := isHappy(2)
	assert.Equal(t, false, r2)
}
