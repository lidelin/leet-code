package _28_odd_even_linked_list

import (
	"github.com/lidelin/leet-code/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOddEvenList(t *testing.T) {
	head := utils.MakeList([]int{1, 2, 3, 4, 5})
	expected := utils.MakeList([]int{1, 3, 5, 2, 4})
	actual := OddEvenList(head)
	assert.Equal(t, expected, actual)

	head = utils.MakeList([]int{2, 1, 3, 5, 6, 4, 7})
	expected = utils.MakeList([]int{2, 3, 6, 7, 1, 5, 4})
	actual = OddEvenList(head)
	assert.Equal(t, expected, actual)
}
