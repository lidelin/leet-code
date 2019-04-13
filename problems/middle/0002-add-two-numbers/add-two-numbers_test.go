package _002_add_two_numbers

import (
	"github.com/stretchr/testify/assert"
	"leet-code/utils"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	list1 := utils.MakeList([]int{2, 4, 3})
	list2 := utils.MakeList([]int{5, 6, 4})

	answer := utils.MakeList([]int{7, 0, 8})

	assert.Equal(t, answer, AddTwoNumbers(list1, list2))
}
