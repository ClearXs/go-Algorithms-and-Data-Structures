package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// 860. 柠檬水找零 https://leetcode.cn/problems/lemonade-change/description/

// 在柠檬水摊上，每一杯柠檬水的售价为 5 美元。顾客排队购买你的产品，（按账单 bills 支付的顺序）一次购买一杯。
//
// 每位顾客只买一杯柠檬水，然后向你付 5 美元、10 美元或 20 美元。你必须给每个顾客正确找零，也就是说净交易是每位顾客向你支付 5 美元。
//
// 注意，一开始你手头没有任何零钱。
//
// 给你一个整数数组 bills ，其中 bills[i] 是第 i 位顾客付的账。如果你能给每位顾客正确找零，返回 true ，否则返回 false 。

func lemonadeChange(bills []int) bool {
	five := 0
	ten := 0
	for _, bill := range bills {
		switch bill {
		case 5:
			five++
		case 10:
			if five < 0 {
				return false
			}
			five--
			ten++
		case 20:
			if ten >= 1 && five >= 1 {
				ten--
				five--
			} else if five >= 3 {
				five -= 3
			} else {
				return false
			}
		}
	}
	return true
}

func TestLemonadeChange(t *testing.T) {
	r1 := lemonadeChange([]int{5, 5, 5, 10, 20})
	assert.Equal(t, true, r1)

	r2 := lemonadeChange([]int{5, 5, 10, 10, 20})
	assert.Equal(t, false, r2)

	r3 := lemonadeChange([]int{5, 5, 10})
	assert.Equal(t, true, r3)
}
