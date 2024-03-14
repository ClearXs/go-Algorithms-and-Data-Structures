package leetcode

import (
	"data-structure/queue"
	"github.com/stretchr/testify/assert"
	"testing"
)

// 347. 前 K 个高频元素 https://leetcode.cn/problems/top-k-frequent-elements/description/
// 给你一个整数数组 nums 和一个整数 k ，请你返回其中出现频率前 k 高的元素。你可以按 任意顺序 返回答案。

type NumFrequent struct {
	Num   int
	Count int
}

func topKFrequent(nums []int, k int) []int {
	frequents := make(map[int]int)
	// 通过map记录每一个元素出现的次数
	for _, num := range nums {
		if frequent, ok := frequents[num]; ok {
			frequents[num] = frequent + 1
		} else {
			frequents[num] = 1
		}
	}
	// 通过最小优先级队列记录每一个元素在队列的位置，通过比较队尾元素与频率k的大小判断是否需要弹出，直至队尾元素大于等于频率k
	priorityQueue := queue.NewMinimumPriorityQueue[NumFrequent](func(t1 *NumFrequent, t2 *NumFrequent) int {
		return t1.Count - t2.Count
	})
	for key, value := range frequents {
		f := NumFrequent{Num: key, Count: value}
		priorityQueue.Push(&f)
	}
	result := make([]int, 0)
	for {
		size := priorityQueue.Size()
		if size == 0 {
			break
		} else if size <= k {
			extreme := priorityQueue.Pop()
			result = append(result, extreme.Num)
		} else {
			priorityQueue.Pop()
		}
	}
	return result
}

func TestTopKFrequent(t *testing.T) {
	num1 := []int{1, 1, 1, 2, 2, 3}
	k1 := 2
	r1 := topKFrequent(num1, k1)
	assert.Equal(t, []int{2, 1}, r1)

	num2 := []int{1}
	k2 := 1
	r2 := topKFrequent(num2, k2)
	assert.Equal(t, []int{1}, r2)
}
