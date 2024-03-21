package leetcode

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func jump(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}
	currentCover := 0
	nextCover := 0
	count := 0
	for i := 0; i <= len(nums)-1; i++ {
		// seek maximum
		// calculate current index add nums[i] for be able to cover distance in current cover
		// why i + nums[i], because first nums[i] is jump to range, and now
		nextCover = int(math.Max(float64(i+nums[i]), float64(nextCover)))
		if i == currentCover {
			count += 1
			currentCover = nextCover
			if currentCover >= len(nums)-1 {
				return count
			}
		}
	}
	return count
}

func TestJump(t *testing.T) {
	r1 := jump([]int{2, 3, 1, 1, 4})
	assert.Equal(t, 2, r1)

	r2 := jump([]int{2, 3, 0, 1, 4})
	assert.Equal(t, 2, r2)

	r3 := jump([]int{1, 2, 1, 1, 1})
	assert.Equal(t, 3, r3)

	r4 := jump([]int{7, 0, 9, 6, 9, 6, 1, 7, 9, 0, 1, 2, 9, 0, 3})
	assert.Equal(t, 2, r4)
}
