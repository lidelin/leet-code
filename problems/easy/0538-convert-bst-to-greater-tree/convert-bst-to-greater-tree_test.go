package _538_convert_bst_to_greater_tree

import (
	"github.com/lidelin/leet-code/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConvertBST(t *testing.T) {
	root := utils.MakeTree([]int{5, 2, 13})
	expected := []int{20, 18, 13}
	actual := utils.InOrderTraversal(convertBST(root))
	assert.Equal(t, expected, actual)
}
