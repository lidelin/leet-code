package _226_invert_binary_tree

import (
	"github.com/lidelin/leet-code/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInvertTree(t *testing.T) {
	root := utils.MakeTree([]int{4, 2, 7, 1, 3, 6, 9})
	expected := []int{4, 7, 9, 6, 2, 3, 1}
	actual := utils.PreOrderTraversal(InvertTree(root))
	assert.Equal(t, expected, actual)
}
