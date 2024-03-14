package leetcode

import (
	"data-structure/list"
	"github.com/stretchr/testify/assert"
	"testing"
)

// 19. 删除链表的倒数第 N 个结点 https://leetcode.cn/problems/remove-nth-node-from-end-of-list/description/

// 给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。

func removeLinkedListNthFromEnd(head *list.ListNode, n int) *list.ListNode {
	dummy := &list.ListNode{Next: head}
	slow, fast := dummy, dummy
	up := n
	// fast 先走nb步
	for fast != nil && up > 0 {
		fast = fast.Next
		up--
	}
	// 走n+1步
	if fast != nil {
		fast = fast.Next
	}
	// slow fast 一起走
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	// 移除
	slow.Next = slow.Next.Next
	return dummy.Next
}

func TestRemoveLinkedListNthFromEnd(t *testing.T) {
	h1 := list.MakeArrayList([]int{1, 2, 3, 4, 5})
	r1 := removeLinkedListNthFromEnd(h1, 2)
	assert.Equal(t, []int{1, 2, 3, 5}, r1.ToArray())

	h2 := list.MakeArrayList([]int{1})
	r2 := removeLinkedListNthFromEnd(h2, 1)
	assert.Nil(t, r2)

	h3 := list.MakeArrayList([]int{1, 2})
	r3 := removeLinkedListNthFromEnd(h3, 1)
	assert.Equal(t, []int{1}, r3.ToArray())
}
