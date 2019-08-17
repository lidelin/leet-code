package _872_leaf_similar_trees

import (
	"github.com/lidelin/leet-code/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLeafSimilar(t *testing.T) {
	root1 := utils.MakeTree([]int{8, 5, 3, 10, 6})
	root2 := utils.MakeTree([]int{9, 4, 6, 3, 10})
	actual := LeafSimilar(root1, root2)
	assert.Equal(t, true, actual)
}
