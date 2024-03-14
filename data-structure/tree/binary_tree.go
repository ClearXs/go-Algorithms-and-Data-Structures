package tree

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

const sentinel = math.MinInt
const empty = 0

func BuildBinaryTree(nums []int) *TreeNode {
	if len(nums) <= 0 {
		return nil
	}
	// 基于某一结点的index，其2 * index = left 2 * index = right的特型构建二叉树
	// 使数组首位插入一个sentinel的结点
	return buildBinaryTreeByArray(append([]int{sentinel}, nums...), 1)
}

func buildBinaryTreeByArray(nums []int, nodeIndex int) *TreeNode {
	if nodeIndex >= len(nums) {
		return nil
	}
	v := nums[nodeIndex]
	if v == empty {
		return nil
	}
	node := &TreeNode{
		Val: v,
	}
	left := buildBinaryTreeByArray(nums, 2*nodeIndex)
	right := buildBinaryTreeByArray(nums, 2*nodeIndex+1)
	node.Left = left
	node.Right = right
	return node
}
