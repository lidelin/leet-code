package queue_stack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNumIslands(t *testing.T) {
	ast := assert.New(t)

	var grid [][]byte
	ast.Equal(0, NumIslands(grid))

	grid = [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}
	ast.Equal(1, NumIslands(grid))

	grid = [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}

	ast.Equal(3, NumIslands(grid))
}
