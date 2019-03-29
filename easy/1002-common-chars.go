package main

import (
	"fmt"
)

func min(x int, y int) int {
	if x > y {
		return y
	}
	return x
}

func commonChars(A []string) []string {
	length := len(A)
	charCounts := make([][26]int, length)

	for i := 0; i < length; i++ {
		for _, char := range A[i] {
			charCounts[i][char-'a']++
		}
	}

	for column := 0; column < 26; column++ {
		minimum := charCounts[0][column]
		for row := 0; row < length; row++ {
			minimum = min(charCounts[row][column], minimum)
		}
		charCounts[0][column] = minimum
	}

	var results []string
	for column := 0; column < 26; column++ {
		for i := 0; i < charCounts[0][column]; i++ {
			results = append(results, string('a'+column))
		}
	}

	return results
}

func main() {
	s := []string{"bella", "label", "roller"}

	fmt.Println(commonChars(s))
}
