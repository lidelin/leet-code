package _257_binary_tree_paths

import (
	"github.com/lidelin/leet-code/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBinaryTreePaths(t *testing.T) {
	root := utils.MakeTree([]int{5, 3, 2, 6})
	expected := []string{"5->3->2", "5->6"}
	actual := binaryTreePaths(root)
	assert.Equal(t, expected, actual)
}

func TestBinaryTreePaths2(t *testing.T) {
	root := utils.MakeTree([]int{5, 3, 2, 6})
	expected := []string{"5->6", "5->3->2"}
	actual := binaryTreePaths2(root)
	assert.Equal(t, expected, actual)
}
