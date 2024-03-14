package leetcode

import (
	"data-structure/tree"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

// 111. 二叉树的最小深度 https://leetcode.cn/problems/minimum-depth-of-binary-tree/description/

// 给定一个二叉树，找出其最小深度。
//
// 最小深度是从根节点到最近叶子节点的最短路径上的节点数量。
//
// 说明：叶子节点是指没有子节点的节点。

const NilDepth = 0

func minDepth(root *tree.TreeNode) int {
	return traversalMinimumDepth(root, 1)
}

func traversalMinimumDepth(node *tree.TreeNode, depth int) int {
	if node == nil {
		return NilDepth
	}
	if isLeaf(node) {
		return depth
	}
	leftPath := traversalMinimumDepth(node.Left, depth+1)
	rightPath := traversalMinimumDepth(node.Right, depth+1)
	if leftPath == NilDepth && rightPath != NilDepth {
		return rightPath
	} else if leftPath != NilDepth && rightPath == NilDepth {
		return leftPath
	} else {
		return int(math.Min(float64(leftPath), float64(rightPath)))
	}
}

func isLeaf(node *tree.TreeNode) bool {
	if node.Left == nil && node.Right == nil {
		return true
	}
	return false
}

func TestMinDepth(t *testing.T) {
	n1 := []int{3, 9, 20, 0, 0, 15, 7}
	t1 := tree.BuildBinaryTree(n1)
	minimum1 := minDepth(t1)
	assert.Equal(t, 2, minimum1)

	//n2 := []int{2, 0, 3, 0, 4, 0, 5, 0, 6}
	//t2 := BuildBinaryTree(n2)
	//minimum2 := minDepth(t2)
	//assert.Equal(t, 5, minimum2)

	n3 := []int{1, 2, 3, 4, 5}
	t3 := tree.BuildBinaryTree(n3)
	minimum3 := minDepth(t3)
	assert.Equal(t, 2, minimum3)
}
