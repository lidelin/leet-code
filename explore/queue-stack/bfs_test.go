package queue_stack

import (
	"github.com/lidelin/leet-code/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBfs(t *testing.T) {
	ast := assert.New(t)

	root := utils.MakeTree([]int{1, 22, 4, 55, 3, 9})

	ast.NotEqual(-1, Bfs(root, 4))
	ast.Equal(-1, Bfs(root, 5))

	ast.NotEqual(-1, Bfs2(root, 4))
	ast.Equal(-1, Bfs2(root, 5))
}
