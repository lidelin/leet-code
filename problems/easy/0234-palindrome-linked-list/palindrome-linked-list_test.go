package _234_palindrome_linked_list

import (
	"github.com/stretchr/testify/assert"
	"leet-code/utils"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	head := utils.MakeList([]int{1, 2, 2, 1})
	actual := IsPalindrome(head)
	assert.Equal(t, true, actual)

	head = utils.MakeList([]int{1, 2, 2, 1, 1})
	actual = IsPalindrome(head)
	assert.Equal(t, false, actual)

	head = utils.MakeList([]int{1, 2, 3, 3, 2, 1})
	actual = IsPalindrome(head)
	assert.Equal(t, true, actual)

	head = utils.MakeList([]int{1, 1})
	actual = IsPalindrome(head)
	assert.Equal(t, true, actual)

	head = utils.MakeList([]int{1, 2})
	actual = IsPalindrome(head)
	assert.Equal(t, false, actual)

	head = utils.MakeList([]int{1})
	actual = IsPalindrome(head)
	assert.Equal(t, true, actual)
}
