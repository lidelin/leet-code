package _654_maximum_binary_tree

import (
	"github.com/lidelin/leet-code/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConstructMaximumBinaryTree(t *testing.T) {
	nums := []int{3, 2, 1, 6, 0, 5}
	expected := []int{6, 3, 2, 1, 5, 0}
	actual := utils.PreOrderTraversal(constructMaximumBinaryTree(nums))
	assert.Equal(t, expected, actual)
}
