package _148_sort_list

import (
	"github.com/stretchr/testify/assert"
	"leet-code/utils"
	"testing"
)

func TestSortListMerge(t *testing.T) {
	head := utils.MakeList([]int{4, 2, 1, 3})
	expected := utils.MakeList([]int{1, 2, 3, 4})
	actual := SortListMerge(head)
	assert.Equal(t, expected, actual)

	head = utils.MakeList([]int{-1, 5, 3, 4, 0})
	expected = utils.MakeList([]int{-1, 0, 3, 4, 5})
	actual = SortListMerge(head)
	assert.Equal(t, expected, actual)
}

func TestSortListQuick(t *testing.T) {
	head := utils.MakeList([]int{4, 2, 1, 3})
	expected := utils.MakeList([]int{1, 2, 3, 4})
	actual := SortListQuick(head)
	assert.Equal(t, expected, actual)

	head = utils.MakeList([]int{-1, 5, 3, 4, 0})
	expected = utils.MakeList([]int{-1, 0, 3, 4, 5})
	actual = SortListQuick(head)
	assert.Equal(t, expected, actual)
}
