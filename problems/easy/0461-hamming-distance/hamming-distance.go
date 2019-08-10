package _461_hamming_distance

import "math/bits"

func HammingDistance(x int, y int) int {
	temp := x ^ y
	count := 0

	for temp != 0 {
		if temp&1 == 1 {
			count++
		}

		temp = temp >> 1
	}

	return count
}

func HammingDistance2(x int, y int) int {
	return bits.OnesCount(uint(x ^ y))
}
