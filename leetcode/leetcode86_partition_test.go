// 86. 分隔链表
// https://leetcode.cn/problems/partition-list/
// 给你一个链表的头节点 head 和一个特定值 x ，请你对链表进行分隔，使得所有 小于 x 的节点都出现在 大于或等于 x 的节点之前。
//
// 你应当 保留 两个分区中每个节点的初始相对位置。
package leetcode

import (
	"data-structure/list"
	"github.com/stretchr/testify/assert"
	"testing"
)

func partition(head *list.ListNode, x int) *list.ListNode {
	p := head
	// 把原题目要求分解为两个链表
	low := &list.ListNode{}
	high := &list.ListNode{}
	// 用于处理边界问题
	v := low
	w := high

	for p != nil {
		if p.Val < x {
			low.Next = p
			low = low.Next
		} else {
			high.Next = p
			high = high.Next
		}
		p = p.Next
	}
	// 去除原有关联关系
	low.Next = nil
	high.Next = nil
	// 合并两个链表
	low.Next = w.Next
	return v.Next
}

func TestPartition(t *testing.T) {
	l1 := list.MakeArrayList([]int{1, 4, 3, 2, 5, 2})
	assert.Equal(t, []int{1, 2, 2, 4, 3, 5}, partition(l1, 3).ToArray())

	l2 := list.MakeArrayList([]int{2, 1})
	assert.Equal(t, []int{1, 2}, partition(l2, 2).ToArray())
}
