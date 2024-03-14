package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// 654. 最大二叉树 https://leetcode.cn/problems/maximum-binary-tree/description/

// 给定一个不重复的整数数组 nums 。 最大二叉树 可以用下面的算法从 nums 递归地构建:
//
// 创建一个根节点，其值为 nums 中的最大值。
// 递归地在最大值 左边 的 子数组前缀上 构建左子树。
// 递归地在最大值 右边 的 子数组后缀上 构建右子树。
// 返回 nums 构建的 最大二叉树 。

type MaximumTreeNode struct {
	Val   int
	Left  *MaximumTreeNode
	Right *MaximumTreeNode
}

func constructMaximumBinaryTree(nums []int) *MaximumTreeNode {
	return maximumBinaryTreeBuild(nums)
}

func maximumBinaryTreeBuild(subArray []int) *MaximumTreeNode {
	maximumIndex := findMaximumIndex(subArray)
	node := MaximumTreeNode{Val: subArray[maximumIndex]}
	if maximumIndex > 0 {
		leftNode := maximumBinaryTreeBuild(subArray[0:maximumIndex])
		node.Left = leftNode
	}
	if maximumIndex < len(subArray)-1 {
		rightNode := maximumBinaryTreeBuild(subArray[maximumIndex+1:])
		node.Right = rightNode
	}
	return &node
}

func findMaximumIndex(subArray []int) int {
	maximumIndex := 0
	for i, v := range subArray {
		if subArray[maximumIndex] < v {
			maximumIndex = i
		}
	}
	return maximumIndex
}

func TestConstructMaximumBinaryTree(t *testing.T) {
	n1 := []int{3, 2, 1}
	node1 := constructMaximumBinaryTree(n1)
	assert.Equal(t, 3, node1.Val)

	n2 := []int{3, 2, 1, 6, 0, 5}
	node2 := constructMaximumBinaryTree(n2)
	assert.Equal(t, 6, node2.Val)
}
