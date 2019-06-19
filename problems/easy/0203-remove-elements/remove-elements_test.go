package _0203_remove_elements

import (
	"github.com/stretchr/testify/assert"
	"leet-code/utils"
	"testing"
)

func TestRemoveElements(t *testing.T) {
	head := utils.MakeList([]int{1, 2, 3, 4, 5, 6})
	expected := utils.MakeList([]int{1, 2, 3, 4, 5})
	actual := RemoveElements(head, 6)
	assert.Equal(t, expected, actual)
}
