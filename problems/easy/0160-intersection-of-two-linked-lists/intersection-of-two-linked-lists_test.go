package _160_intersection_of_two_linked_lists

import (
	"github.com/lidelin/leet-code/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetIntersectionNode(t *testing.T) {
	intersection := utils.MakeList([]int{8, 4, 5})
	l1 := utils.MakeList([]int{4, 1})
	l2 := utils.MakeList([]int{5, 0, 1})
	utils.CatList(l1, intersection)
	utils.CatList(l2, intersection)
	actual := GetIntersectionNode(l1, l2)
	assert.Equal(t, intersection, actual)

	intersection = utils.MakeList([]int{2, 4})
	l1 = utils.MakeList([]int{0, 9, 1})
	l2 = utils.MakeList([]int{3})
	utils.CatList(l1, intersection)
	utils.CatList(l2, intersection)
	actual = GetIntersectionNode(l1, l2)
	assert.Equal(t, intersection, actual)

	var expected *ListNode = nil
	l1 = utils.MakeList([]int{2, 6, 4})
	l2 = utils.MakeList([]int{1, 5})
	actual = GetIntersectionNode(l1, l2)
	assert.Equal(t, expected, actual)
}
