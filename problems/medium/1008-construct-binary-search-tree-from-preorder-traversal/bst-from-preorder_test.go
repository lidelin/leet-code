package _008_construct_binary_search_tree_from_preorder_traversal

import (
	"github.com/lidelin/leet-code/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBstFromPreorder(t *testing.T) {
	preorder := []int{8, 5, 1, 7, 10, 12}
	root := bstFromPreorder(preorder)
	actual := utils.PreOrderTraversal(root)
	assert.Equal(t, preorder, actual)
}
