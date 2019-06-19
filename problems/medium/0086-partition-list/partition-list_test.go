package _086_partition_list

import (
	"github.com/stretchr/testify/assert"
	"leet-code/utils"
	"testing"
)

func TestPartition(t *testing.T) {
	head := utils.MakeList([]int{1, 4, 3, 2, 5, 2})
	x := 3
	expected := utils.MakeList([]int{1, 2, 2, 4, 3, 5})
	actual := Partition(head, x)
	assert.Equal(t, expected, actual)

	head = utils.MakeList([]int{1, 1})
	x = 0
	expected = utils.MakeList([]int{1, 1})
	actual = Partition(head, x)
	assert.Equal(t, expected, actual)

	head = utils.MakeList([]int{2, 1})
	x = 2
	expected = utils.MakeList([]int{1, 2})
	actual = Partition(head, x)
	assert.Equal(t, expected, actual)
}
