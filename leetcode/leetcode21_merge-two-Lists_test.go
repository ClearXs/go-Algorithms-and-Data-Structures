// 21. 合并两个有序链表
// https://leetcode.cn/problems/merge-two-sorted-lists/
// 将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
package leetcode

import (
	"data-structure/list"
	"github.com/stretchr/testify/assert"
	"testing"
)

// 注意这题是两个链表合并称新链表，通过双指针无法做出，需要多一个指针指向新的链表
func mergeTwoLists(list1 *list.ListNode, list2 *list.ListNode) *list.ListNode {
	// 虚拟头节点，可以用来处理单链表边界问题
	virtual := &list.ListNode{}
	p := virtual
	// 以l1为基准比较
	l1, l2 := list1, list2
	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			p.Next = l2
			l2 = l2.Next
		} else {
			p.Next = l1
			l1 = l1.Next
		}
		p = p.Next
	}

	if l1 == nil {
		p.Next = l2
	} else if l2 == nil {
		p.Next = l1
	}
	return virtual.Next
}

func TestMergeTwoLists(t *testing.T) {
	l1 := list.MakeArrayList([]int{1, 2, 4})
	l2 := list.MakeArrayList([]int{1, 3, 4})

	m1 := mergeTwoLists(l1, l2)
	assert.Equal(t, []int{1, 1, 2, 3, 4, 4}, m1.ToArray())

	l3 := list.MakeArrayList([]int{})
	l4 := list.MakeArrayList([]int{})
	m2 := mergeTwoLists(l3, l4)
	assert.Equal(t, []int{}, m2.ToArray())

	l5 := list.MakeArrayList([]int{})
	l6 := list.MakeArrayList([]int{0})
	m3 := mergeTwoLists(l5, l6)
	assert.Equal(t, []int{0}, m3.ToArray())
}
