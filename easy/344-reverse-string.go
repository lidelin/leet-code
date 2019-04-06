package main

import "fmt"

func reverseString(s []byte) {
	length := len(s)
	if length <= 0 {
		return
	}

	tail := length - 1
	loop := length / 2
	temp := s[0]

	for head := 0; head < loop; head++ {
		temp = s[head]
		s[head] = s[tail]
		s[tail] = temp
		tail--
	}
}

func main() {
	str := []byte{'h', 'e', 'l', 'l', 'o', 'o'}
	fmt.Println(string(str))

	reverseString(str)
	fmt.Println(string(str))
}
