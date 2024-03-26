package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// 509. 斐波那契数 https://leetcode.cn/problems/fibonacci-number/description/

// 斐波那契数 （通常用 F(n) 表示）形成的序列称为 斐波那契数列 。该数列由 0 和 1 开始，后面的每一项数字都是前面两项数字的和。也就是：
//
// F(0) = 0，F(1) = 1
// F(n) = F(n - 1) + F(n - 2)，其中 n > 1
// 给定 n ，请计算 F(n) 。

func fib(n int) int {
	dp0, dp1, dpN := 0, 1, 0
	for i := 0; i < n; i++ {
		if i > 1 {
			dp0 = dp1
			dp1 = dpN
		}
		dpN = dp0 + dp1
	}
	return dpN
}

func TestFib(t *testing.T) {
	r1 := fib(2)
	assert.Equal(t, 1, r1)

	r2 := fib(3)
	assert.Equal(t, 2, r2)

	r3 := fib(4)
	assert.Equal(t, 3, r3)
}
