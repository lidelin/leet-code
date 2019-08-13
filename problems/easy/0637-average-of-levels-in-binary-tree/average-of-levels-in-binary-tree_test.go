package _637_average_of_levels_in_binary_tree

import (
	"github.com/lidelin/leet-code/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAverageOfLevels(t *testing.T) {
	root := utils.MakeTree([]int{3, 2, 20, 15, 17})
	expected := []float64{3, 11, 15, 17}
	actual := AverageOfLevels(root)
	assert.Equal(t, expected, actual)
}
