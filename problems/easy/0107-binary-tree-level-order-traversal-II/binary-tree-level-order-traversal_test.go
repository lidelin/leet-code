package _107_binary_tree_level_order_traversal_II

import (
	"github.com/lidelin/leet-code/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLevelOrderBottom(t *testing.T) {
	root := utils.MakeTree([]int{3, 2, 20, 15, 17})
	expected := [][]int{{17}, {15}, {2, 20}, {3}}
	actual := LevelOrderBottom(root)
	assert.Equal(t, expected, actual)
}
