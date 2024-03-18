package leetcode

import (
	"github.com/stretchr/testify/assert"
	"sort3"
	"testing"
)

func findContentChildren(g []int, s []int) int {
	g = sort3.QuickSort(g)
	s = sort3.QuickSort(s)

	count := 0
	maxCookiesIndex := len(s) - 1

	for i := len(g) - 1; i > -1; i-- {
		if maxCookiesIndex < 0 {
			break
		}
		appetite := g[i]
		cookiesIndex := findContentCookiesIndex(s, appetite, maxCookiesIndex)
		if cookiesIndex > -1 {
			count += 1
			// remove consume cookies
			maxCookiesIndex = cookiesIndex - 1
		}
	}
	return count
}

func findContentCookiesIndex(s []int, appetite int, lastCookiesIndex int) int {
	for i := lastCookiesIndex; i > -1; i-- {
		if s[i] >= appetite {
			return i
		}
	}
	return -1
}

func TestFindContentChildren(t *testing.T) {
	r1 := findContentChildren([]int{1, 2, 3}, []int{1, 1})
	assert.Equal(t, 1, r1)

	r2 := findContentChildren([]int{1, 2}, []int{1, 2, 3})
	assert.Equal(t, 2, r2)

	r3 := findContentChildren([]int{1, 2, 3}, []int{3})
	assert.Equal(t, 1, r3)
}
