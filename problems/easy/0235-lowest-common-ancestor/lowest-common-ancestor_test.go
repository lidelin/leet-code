package _235_lowest_common_ancestor

import (
	"github.com/lidelin/leet-code/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLowestCommonAncestor(t *testing.T) {
	root := utils.MakeTree([]int{6, 2, 8, 0, 4, 7, 9, 3, 5})
	p := &TreeNode{Val: 2}
	q := &TreeNode{Val: 8}
	actual := LowestCommonAncestor(root, p, q)
	assert.Equal(t, 6, actual.Val)

	p = &TreeNode{Val: 2}
	q = &TreeNode{Val: 4}
	actual = LowestCommonAncestor(root, p, q)
	assert.Equal(t, 2, actual.Val)
}

func TestLowestCommonAncestorIteration(t *testing.T) {
	root := utils.MakeTree([]int{6, 2, 8, 0, 4, 7, 9, 3, 5})
	p := &TreeNode{Val: 2}
	q := &TreeNode{Val: 8}
	actual := LowestCommonAncestorIteration(root, p, q)
	assert.Equal(t, 6, actual.Val)

	p = &TreeNode{Val: 2}
	q = &TreeNode{Val: 4}
	actual = LowestCommonAncestorIteration(root, p, q)
	assert.Equal(t, 2, actual.Val)
}
