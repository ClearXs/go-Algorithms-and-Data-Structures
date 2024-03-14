// https://leetcode.cn/problems/linked-list-cycle/
// 环形链表
// 给你一个链表的头节点 head ，判断链表中是否有环。
//
// 如果链表中有某个节点，可以通过连续跟踪 next 指针再次到达，则链表中存在环。 为了表示给定链表中的环，评测系统内部使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。注意：pos 不作为参数进行传递 。仅仅是为了标识链表的实际情况。
//
// 如果链表中存在环 ，则返回 true 。 否则，返回 false 。
package leetcode

import (
	"data-structure/list"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCycle(t *testing.T) {
	l1 := list.ListNode{Val: 1}
	l2 := list.ListNode{Val: 2}
	l1.Next = &l2
	l2.Next = &l1
	assert.Equal(t, true, hasCycle(&l1))
	assert.Equal(t, true, hasCycleSlowAndQuick(&l1))
}

func TestSingleCycle(t *testing.T) {
	l1 := list.ListNode{Val: 1}
	assert.Equal(t, false, hasCycle(&l1))
}

func TestMultiCycle(t *testing.T) {
	l1 := list.ListNode{Val: 3}
	l2 := list.ListNode{Val: 2}
	l3 := list.ListNode{}
	l4 := list.ListNode{Val: -4}
	l1.Next = &l2
	l2.Next = &l3
	l3.Next = &l4
	l4.Next = &l2
	assert.Equal(t, true, hasCycle(&l1))
}

// 通过使用额外空间
// 使用map结构，如果是环则环的元素一定是存储在map里面
func hasCycle(head *list.ListNode) bool {
	// map链接方法
	n := head
	hash := make(map[*list.ListNode]int)
	for n != nil {
		_, o := hash[n]
		if o {
			return true
		}
		hash[n] = n.Val
		n = n.Next
	}
	return false
}

// 快慢指针法
func hasCycleSlowAndQuick(head *list.ListNode) bool {
	low, fast := head, head
	for fast != nil && fast.Next != nil {
		low = low.Next
		fast = fast.Next.Next
		if low == fast {
			return true
		}
	}
	return false
}

func TestHasCycleSlowAndQuick(t *testing.T) {
	l1 := list.MakeArrayList([]int{3, 2, 0, -4})
	assert.Equal(t, true, hasCycleSlowAndQuick(l1))
}
