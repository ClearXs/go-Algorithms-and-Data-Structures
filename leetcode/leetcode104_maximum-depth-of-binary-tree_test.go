package leetcode

import (
	"data-structure/tree"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

// 104. 二叉树的最大深度 https://leetcode.cn/problems/maximum-depth-of-binary-tree/description/

// 给定一个二叉树 root ，返回其最大深度。
//
// 二叉树的 最大深度 是指从根节点到最远叶子节点的最长路径上的节点数。

const MaximumNilDepth = 0

func maxDepth(root *tree.TreeNode) int {
	return traversalMaximumDepth(root, 1)
}

func traversalMaximumDepth(node *tree.TreeNode, depth int) int {
	if node == nil {
		return MaximumNilDepth
	}
	if isMaximumLeaf(node) {
		return depth
	}
	leftPath := traversalMaximumDepth(node.Left, depth+1)
	rightPath := traversalMaximumDepth(node.Right, depth+1)
	if leftPath == MaximumNilDepth && rightPath != MaximumNilDepth {
		return rightPath
	} else if leftPath != MaximumNilDepth && rightPath == MaximumNilDepth {
		return leftPath
	} else {
		return int(math.Max(float64(leftPath), float64(rightPath)))
	}
}

func isMaximumLeaf(node *tree.TreeNode) bool {
	if node.Left == nil && node.Right == nil {
		return true
	}
	return false
}

func TestMaxDepth(t *testing.T) {
	n1 := []int{3, 9, 20, 0, 0, 15, 7}
	t1 := tree.BuildBinaryTree(n1)
	minimum1 := maxDepth(t1)
	assert.Equal(t, 3, minimum1)

	//n2 := []int{2, 0, 3, 0, 4, 0, 5, 0, 6}
	//t2 := BuildBinaryTree(n2)
	//minimum2 := minDepth(t2)
	//assert.Equal(t, 5, minimum2)

	n3 := []int{1, 2, 3, 4, 5}
	t3 := tree.BuildBinaryTree(n3)
	minimum3 := maxDepth(t3)
	assert.Equal(t, 3, minimum3)
}
