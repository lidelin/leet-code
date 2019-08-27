package _100_same_tree

import (
	"github.com/lidelin/leet-code/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsSameTree(t *testing.T) {
	p := utils.MakeTree([]int{1, 2, 3})
	q := utils.MakeTree([]int{1, 2, 3})
	actual := isSameTree(p, q)
	assert.Equal(t, true, actual)

	p = utils.MakeTree([]int{1, 2})
	q = utils.MakeTree([]int{2, 1})
	actual = isSameTree(p, q)
	assert.Equal(t, false, actual)
}
