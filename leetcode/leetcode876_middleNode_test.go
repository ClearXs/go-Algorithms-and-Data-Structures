// 876. 链表的中间结点
// https://leetcode.cn/problems/middle-of-the-linked-list/
// 给你单链表的头结点 head ，请你找出并返回链表的中间结点。
//
// 如果有两个中间结点，则返回第二个中间结点。
package leetcode

import (
	"data-structure/list"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMiddleNode(t *testing.T) {
	l1 := list.MakeArrayList([]int{1, 2, 3, 4, 5})
	assert.Equal(t, []int{3, 4, 5}, middleNode(l1).ToArray())

	l2 := list.MakeArrayList([]int{1, 2, 3, 4, 5, 6})
	assert.Equal(t, []int{4, 5, 6}, middleNode(l2).ToArray())
}

// 快慢指针法
// 假设节点长度为n，那么：
// 令每一次走k步，则快指针走2k，则当2k=n时，慢指针恰好处于终点位置
func middleNode(head *list.ListNode) *list.ListNode {
	low, fast := head, head
	for fast != nil && fast.Next != nil {
		low = low.Next
		fast = fast.Next.Next
	}
	return low
}
