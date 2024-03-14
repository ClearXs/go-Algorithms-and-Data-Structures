package leetcode

import (
	"data-structure/list"
	"github.com/stretchr/testify/assert"
	"testing"
)

// 203. 移除链表元素 https://leetcode.cn/problems/remove-linked-list-elements/description/

// 给你一个链表的头节点 head 和一个整数 val ，请你删除链表中所有满足 Node.val == val 的节点，并返回 新的头节点 。

func removeLinkedListElements(head *list.ListNode, val int) *list.ListNode {
	if head == nil {
		return nil
	}
	dummy := &list.ListNode{Next: head}
	pre, p := dummy, dummy

	for p != nil {
		if p.Val == val {
			pre.Next = p.Next
		} else {
			pre = p
		}
		p = p.Next
	}
	return dummy.Next

}

func TestRemoveElements(t *testing.T) {
	h1 := list.MakeArrayList([]int{1, 2, 6, 3, 4, 5, 6})
	r1 := removeLinkedListElements(h1, 6)
	assert.Equal(t, []int{1, 2, 3, 4, 5}, r1.ToArray())

	h2 := list.MakeArrayList([]int{})
	r2 := removeLinkedListElements(h2, 1)
	assert.Nil(t, r2)

	h3 := list.MakeArrayList([]int{7, 7, 7, 7})
	r3 := removeLinkedListElements(h3, 7)
	assert.Nil(t, r3)
}
