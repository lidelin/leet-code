package _669_trim_a_binary_search_tree

import (
	"github.com/lidelin/leet-code/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTrimBST(t *testing.T) {
	root := utils.MakeTree([]int{1, 0, 2})
	expected := []int{1, 2}
	actual := utils.PreOrderTraversal(TrimBST(root, 1, 2))
	assert.Equal(t, expected, actual)
}
