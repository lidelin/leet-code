package main

import "fmt"

func numJewelsInStones(J string, S string) int {
	jewelsMap := make(map[int32]int)

	for _, value := range J {
		jewelsMap[value] = 1
	}

	count := 0

	for _, stone := range S {
		if _, exists := jewelsMap[stone]; exists {
			count++
		}
	}

	return count
}

func main() {
	fmt.Println(numJewelsInStones("aA", "aAAbbbb"))
}
