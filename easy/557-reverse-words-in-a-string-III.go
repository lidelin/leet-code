package main

import "fmt"

func reverseWords(s string) string {
	length := len(s)
	newString := make([]byte, length)
	blank := []byte(" ")[0]
	head, tail := 0, 0

	for current := 0; current <= length; current++ {
		if current == length || s[current] == blank {
			tail = current - 1
			for {
				if head >= current {
					break
				}

				newString[head] = s[tail]
				head++
				tail--
			}
			if current < length {
				newString[head] = s[current]
				head++
				fmt.Println(string(newString))
			}
		}
	}

	return string(newString)
}

func main() {
	fmt.Println(reverseWords("Let's take LeetCode contest"))
}
