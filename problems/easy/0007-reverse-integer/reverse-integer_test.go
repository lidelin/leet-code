package _0007_reverse_integer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverse(t *testing.T) {
	temp := 123
	actual := Reverse(temp)
	expected := 321
	assert.Equal(t, expected, actual)

	temp = 1534236469
	actual = Reverse(temp)
	expected = 0
	assert.Equal(t, expected, actual)
}
