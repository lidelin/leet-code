package main

import "fmt"

func subsets(nums []int) [][]int {
	results := [][]int{{}}

	for _, num := range nums {
		for _, result := range results {
			length := len(result) + 1
			temp := make([]int, length)
			copy(temp, result)
			temp[length-1] = num
			results = append(results, temp)
		}
	}

	return results
}

func main() {
	fmt.Println(subsets([]int{9, 0, 3, 5, 7}))
}
