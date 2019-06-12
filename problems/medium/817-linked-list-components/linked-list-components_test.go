package _817_linked_list_components

import (
	"github.com/stretchr/testify/assert"
	"leet-code/utils"
	"testing"
)

func TestNumComponents(t *testing.T) {
	head := utils.MakeList([]int{0, 1, 2, 3})
	g := []int{0, 1, 3}
	expected := 2
	actual := NumComponents(head, g)
	assert.Equal(t, expected, actual)

	head = utils.MakeList([]int{0, 1, 2, 3, 4})
	g = []int{0, 3, 1, 4}
	expected = 2
	actual = NumComponents(head, g)
	assert.Equal(t, expected, actual)

	head = utils.MakeList([]int{0, 1, 2})
	g = []int{1, 0}
	expected = 1
	actual = NumComponents(head, g)
	assert.Equal(t, expected, actual)
}
