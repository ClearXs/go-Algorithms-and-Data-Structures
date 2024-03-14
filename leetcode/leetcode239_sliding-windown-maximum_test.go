package leetcode

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

// 239. 滑动窗口最大值 https://leetcode.cn/problems/sliding-window-maximum/description/
// 给你一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。
//
// 返回 滑动窗口中的最大值 。

const NullInt = math.MaxInt

// 该题的题解思路是使用单调双端队列
// 当窗口移动后，便可以弹出单调最后的首位元素（该元素一定是最大值）
func maxSlidingWindow(nums []int, k int) []int {
	count := 0
	deque := NewMonotonicDeque()
	maximum := make([]int, 0)
	for index, num := range nums {
		if index >= k {
			// 当滑动窗口移动时，此时移除最先元素
			deque.Pop(nums[index-k])
		}
		deque.Push(num)
		if count != k {
			count += 1
		}

		if count == k {
			extremum := deque.Front()
			if extremum != math.MaxInt {
				maximum = append(maximum, extremum)
			}
		}
	}
	return maximum
}

// 单调双端队列
type MonotonicDeque struct {
	items []int
}

func NewMonotonicDeque() *MonotonicDeque {
	return &MonotonicDeque{items: make([]int, 0)}
}

// UnsafePop 直接弹出元素
func (deque *MonotonicDeque) UnsafePop() int {
	extremum := deque.Front()
	deque.items = deque.items[1:]
	return extremum
}

// Pop 如果给定的Pop元素是末尾元素（即极值元素）那么弹出，否则不弹出
func (deque *MonotonicDeque) Pop(value int) int {
	extremum := deque.Front()
	if extremum == NullInt {
		return NullInt
	}
	if extremum == value {
		deque.items = deque.items[1:]
		return extremum
	}
	return value
}

// Push 如果队列push的元素大于队列首位元素，那么移除首位元素，直至没有大于。
func (deque *MonotonicDeque) Push(value int) {
	back := deque.Back()
	if back == NullInt {
		deque.items = append(deque.items, value)
	} else {
		for back != NullInt && value > back {
			deque.items = deque.items[0 : deque.Size()-1]
			back = deque.Back()
		}
		deque.items = append(deque.items, value)
	}
}

// Front 返回队列首位元素
func (deque *MonotonicDeque) Back() int {
	if deque.Size() == 0 {
		return NullInt
	}
	return deque.items[deque.Size()-1]
}

// Back 返回队列出口元素
func (deque *MonotonicDeque) Front() int {
	if deque.Size() == 0 {
		return NullInt
	}
	return deque.items[0]
}

// Size 返回队列大小
func (deque *MonotonicDeque) Size() int {
	return len(deque.items)
}

func TestMaxSlidingWindow(t *testing.T) {
	nums1 := []int{1, 3, -1, -3, 5, 3, 6, 7}
	maximum1 := maxSlidingWindow(nums1, 3)
	assert.Equal(t, []int{3, 3, 5, 5, 6, 7}, maximum1)

	nums2 := []int{1}
	maximum2 := maxSlidingWindow(nums2, 1)
	assert.Equal(t, []int{1}, maximum2)

	nums3 := []int{1, -1}
	maximum3 := maxSlidingWindow(nums3, 1)
	assert.Equal(t, []int{1, -1}, maximum3)

	nums4 := []int{7, 2, 4}
	maximum4 := maxSlidingWindow(nums4, 2)
	assert.Equal(t, []int{7, 4}, maximum4)

	nums5 := []int{1, 3, 1, 2, 0, 5}
	maximum5 := maxSlidingWindow(nums5, 3)
	assert.Equal(t, []int{3, 3, 2, 5}, maximum5)

	nums6 := []int{-7, -8, 7, 5, 7, 1, 6, 0}
	maximum6 := maxSlidingWindow(nums6, 4)
	assert.Equal(t, []int{7, 7, 7, 7, 7}, maximum6)
}
