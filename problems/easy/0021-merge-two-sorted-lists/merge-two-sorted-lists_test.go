package _021_merge_two_sorted_lists

import (
	"github.com/stretchr/testify/assert"
	"leet-code/utils"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	l1 := utils.MakeList([]int{1, 2, 4})
	l2 := utils.MakeList([]int{1, 3, 4})
	expected := utils.MakeList([]int{1, 1, 2, 3, 4, 4})
	actual := MergeTwoLists(l1, l2)
	assert.Equal(t, expected, actual)

	l1 = utils.MakeList([]int{1, 2, 4, 5, 6})
	l2 = utils.MakeList([]int{1, 3, 4})
	expected = utils.MakeList([]int{1, 1, 2, 3, 4, 4, 5, 6})
	actual = MergeTwoLists(l1, l2)
	assert.Equal(t, expected, actual)
}

func TestMergeTwoListsRecursively(t *testing.T) {
	l1 := utils.MakeList([]int{1, 2, 4})
	l2 := utils.MakeList([]int{1, 3, 4})
	expected := utils.MakeList([]int{1, 1, 2, 3, 4, 4})
	actual := MergeTwoListsRecursively(l1, l2)
	assert.Equal(t, expected, actual)

	l1 = utils.MakeList([]int{1, 2, 4, 5, 6})
	l2 = utils.MakeList([]int{1, 3, 4})
	expected = utils.MakeList([]int{1, 1, 2, 3, 4, 4, 5, 6})
	actual = MergeTwoListsRecursively(l1, l2)
	assert.Equal(t, expected, actual)
}
