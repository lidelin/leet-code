package queue_stack

import (
	"github.com/lidelin/leet-code/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDfs(t *testing.T) {
	root := utils.MakeTree([]int{3, 8, 1, 6, 2, 5})
	actual := Dfs(root, 2)
	assert.Equal(t, true, actual)

	actual = Dfs(root, 20)
	assert.Equal(t, false, actual)
}

func TestDfs2(t *testing.T) {
	root := utils.MakeTree([]int{3, 8, 1, 6, 2, 5})
	actual := Dfs2(root, 2)
	assert.Equal(t, true, actual)

	actual = Dfs2(root, 20)
	assert.Equal(t, false, actual)
}
