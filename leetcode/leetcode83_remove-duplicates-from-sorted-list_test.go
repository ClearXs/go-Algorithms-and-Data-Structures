// 83. 删除排序链表中的重复元素
// https://leetcode.cn/problems/remove-duplicates-from-sorted-list/
// 给定一个已排序的链表的头 head ， 删除所有重复的元素，使每个元素只出现一次 。返回 已排序的链表 。
package leetcode

import (
	"data-structure/list"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDeleteDuplicates1(t *testing.T) {
	h1 := &list.ListNode{Val: 1}
	h2 := &list.ListNode{Val: 1}
	h3 := &list.ListNode{Val: 2}
	h1.Next = h2
	h2.Next = h3

	ex1 := h1
	ex2 := h3
	ex1.Next = ex2
	// 去重结果
	dd := deleteDuplicates(h1)

	assert.Equal(t, ex1, dd)
}

func TestDeleteDuplicates2(t *testing.T) {
	h1 := &list.ListNode{Val: 1}
	h3 := &list.ListNode{Val: 1}
	h1.Next = h3
	ex1 := h1
	ex1.Next = nil
	// 去重结果
	dd := deleteDuplicates(h1)

	assert.Equal(t, ex1, dd)
}

func deleteDuplicates(head *list.ListNode) *list.ListNode {
	// 需要注意判定边界条件
	if head == nil {
		return head
	}
	slow, fast := head, head
	for fast != nil {
		if slow.Val != fast.Val {
			slow.Next = fast
			slow = slow.Next
		}
		fast = fast.Next
	}
	slow.Next = nil
	return head
}
