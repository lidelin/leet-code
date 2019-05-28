package _147_insertion_sort_list

import (
	"github.com/stretchr/testify/assert"
	"leet-code/utils"
	"testing"
)

func TestInsertionSortList(t *testing.T) {
	head := utils.MakeList([]int{4, 2, 1, 3})
	expected := utils.MakeList([]int{1, 2, 3, 4})
	actual := InsertionSortList(head)
	assert.Equal(t, expected, actual)

	head = utils.MakeList([]int{1, 2, 3, 4})
	expected = utils.MakeList([]int{1, 2, 3, 4})
	actual = InsertionSortList(head)
	assert.Equal(t, expected, actual)

	head = utils.MakeList([]int{2, 1})
	expected = utils.MakeList([]int{1, 2})
	actual = InsertionSortList(head)
	assert.Equal(t, expected, actual)

	head = utils.MakeList([]int{1})
	expected = utils.MakeList([]int{1})
	actual = InsertionSortList(head)
	assert.Equal(t, expected, actual)
}
