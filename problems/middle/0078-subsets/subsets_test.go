package _078_subsets

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSubsets(t *testing.T) {
	input := []int{9, 0, 3, 5, 7}
	output := Subsets(input)

	expected := [][]int{
		{},

		{9},

		{0},
		{9, 0},

		{3},
		{9, 3},
		{0, 3},
		{9, 0, 3},

		{5},
		{9, 5},
		{0, 5},
		{9, 0, 5},
		{3, 5},
		{9, 3, 5},
		{0, 3, 5},
		{9, 0, 3, 5},

		{7},
		{9, 7},
		{0, 7},
		{9, 0, 7},
		{3, 7},
		{9, 3, 7},
		{0, 3, 7},
		{9, 0, 3, 7},
		{5, 7},
		{9, 5, 7},
		{0, 5, 7},
		{9, 0, 5, 7},
		{3, 5, 7},
		{9, 3, 5, 7},
		{0, 3, 5, 7},
		{9, 0, 3, 5, 7},
	}

	assert.Equal(t, expected, output)
}
