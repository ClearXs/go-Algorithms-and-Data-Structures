package leetcode

import (
	"data-structure/tree"
	"github.com/stretchr/testify/assert"
	"testing"
)

// 101. 对称二叉树 https://leetcode.cn/problems/symmetric-tree/description/

// 给你一个二叉树的根节点 root ， 检查它是否轴对称。

func isSymmetric(root *tree.TreeNode) bool {
	if root == nil {
		return false
	}
	if root.Left != nil && root.Right != nil {
		return root.Left.Val == root.Right.Val &&
			detectionSymmetric(root.Left, root.Right)
	}
	return compareSymmetric(root.Left, root.Right)
}

func detectionSymmetric(node *tree.TreeNode, sibling *tree.TreeNode) bool {
	symmetrical := true
	symmetrical = symmetrical && compareSymmetric(node.Left, sibling.Right)
	if symmetrical {
		symmetrical = symmetrical && compareSymmetric(node.Right, sibling.Left)
	}
	if !symmetrical {
		return false
	}
	if node.Left != nil && sibling.Right != nil {
		symmetrical = detectionSymmetric(node.Left, sibling.Right)
		if !symmetrical {
			return false
		}
	}
	if node.Right != nil && sibling.Left != nil {
		symmetrical = detectionSymmetric(node.Right, sibling.Left)
		if !symmetrical {
			return false
		}
	}
	return symmetrical
}

func compareSymmetric(node *tree.TreeNode, sibling *tree.TreeNode) bool {
	if node == nil && sibling == nil {
		return true
	} else if node == nil || sibling == nil {
		return false
	} else {
		return node.Val == sibling.Val
	}
}

func TestIsSymmetric(t *testing.T) {
	r1 := tree.BuildBinaryTree([]int{1, 2, 2, 3, 4, 4, 3})
	s1 := isSymmetric(r1)
	assert.Equal(t, true, s1)

	r2 := tree.BuildBinaryTree([]int{1, 2, 2, 0, 3, 0, 3})
	s2 := isSymmetric(r2)
	assert.Equal(t, false, s2)

	r3 := tree.BuildBinaryTree([]int{1, 0})
	s3 := isSymmetric(r3)
	assert.Equal(t, true, s3)
}
