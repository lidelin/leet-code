package _461_hamming_distance

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHammingDistance(t *testing.T) {
	x := 1
	y := 4
	expected := 2
	actual := HammingDistance(x, y)
	assert.Equal(t, expected, actual)
}
