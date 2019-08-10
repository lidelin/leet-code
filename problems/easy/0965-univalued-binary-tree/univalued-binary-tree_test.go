package _965_univalued_binary_tree

import (
	"github.com/lidelin/leet-code/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsUnivalTree(t *testing.T) {
	root := utils.MakeTree([]int{1, 2, 3, 4, 5, 6})
	actual := IsUnivalTree(root)
	assert.Equal(t, false, actual)

	root = utils.MakeTree([]int{1, 1, 1, 1, 1, 1, 1})
	actual = IsUnivalTree(root)
	assert.Equal(t, true, actual)
}
