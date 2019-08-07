package _700_search_in_a_binary_search_tree

import (
	"github.com/lidelin/leet-code/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSearchBST(t *testing.T) {
	root := utils.MakeTree([]int{4, 2, 7, 1, 3})
	expected := utils.MakeTree([]int{2, 1, 3})
	actual := SearchBST(root, 2)

	assert.Equal(t, expected, actual)
}

func TestSearchBST2(t *testing.T) {
	root := utils.MakeTree([]int{4, 2, 7, 1, 3})
	expected := utils.MakeTree([]int{2, 1, 3})
	actual := SearchBST2(root, 2)

	assert.Equal(t, expected, actual)
}
