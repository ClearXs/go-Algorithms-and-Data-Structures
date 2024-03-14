package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// 383. 赎金信 https://leetcode.cn/problems/ransom-note/description/

// 给你两个字符串：ransomNote 和 magazine ，判断 ransomNote 能不能由 magazine 里面的字符构成。
//
// 如果可以，返回 true ；否则返回 false 。
//
// magazine 中的每个字符只能在 ransomNote 中使用一次。

func canConstruct(ransomNote string, magazine string) bool {
	magazineCharsMap := make(map[int32]int)
	for _, b := range magazine {
		if count, ok := magazineCharsMap[b]; ok {
			magazineCharsMap[b] = count + 1
		} else {
			magazineCharsMap[b] = 1
		}
	}
	for _, b := range ransomNote {
		count, ok := magazineCharsMap[b]
		if !ok {
			return false
		}
		if count >= 1 {
			count = count - 1
			if count == 0 {
				// 把值移除
				delete(magazineCharsMap, b)
			} else {
				magazineCharsMap[b] = count

			}
		}
	}
	return true
}

func TestCanConstruct(t *testing.T) {
	r1 := canConstruct("a", "b")
	assert.Equal(t, false, r1)

	r2 := canConstruct("aa", "ab")
	assert.Equal(t, false, r2)

	r3 := canConstruct("aa", "aab")
	assert.Equal(t, true, r3)
}
