package list

type ListNode struct {
	Val  int
	Next *ListNode
}

// ToArray 返回数组
func (l *ListNode) ToArray() []int {
	nums := make([]int, 0)
	p := l
	for p != nil {
		nums = append(nums, p.Val)
		p = p.Next
	}
	return nums
}

func MakeArrayList(nums []int) *ListNode {
	v := &ListNode{}
	p := v
	for _, ele := range nums {
		p.Next = &ListNode{Val: ele}
		p = p.Next
	}
	return v.Next
}
