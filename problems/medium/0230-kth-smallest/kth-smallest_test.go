package _0230_kth_smallest

import (
	"github.com/stretchr/testify/assert"
	"leet-code/utils"
	"testing"
)

func TestKthSmallest(t *testing.T) {
	root := utils.MakeTree([]int{1, 2, 3, 4, 5, 6})
	expected := 5
	actual := KthSmallest(root, 5)
	assert.Equal(t, expected, actual)
}
