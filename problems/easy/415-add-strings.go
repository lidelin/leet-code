package main

import "fmt"

func byte2int(num byte) int {
	return int(num - '0')
}

func int2byte(num int) byte {
	return byte(num + '0')
}

func reverseString(str []byte) string {
	var result []byte

	length := len(str)
	for i := 0; i < length; i++ {
		result = append(result, str[length-i-1])
	}

	return string(result)
}

func addStrings(num1 string, num2 string) string {
	num1Length := len(num1) - 1
	num2Length := len(num2) - 1

	var carry = 0
	var result []byte

	for {
		if num1Length < 0 && num2Length < 0 {
			break
		}

		if num1Length >= 0 {
			carry += byte2int(num1[num1Length])
			num1Length--
		}

		if num2Length >= 0 {
			carry += byte2int(num2[num2Length])
			num2Length--
		}

		result = append(result, int2byte(carry%10))
		carry /= 10
	}

	if carry > 0 {
		result = append(result, int2byte(carry))
	}

	return reverseString(result)
}

func main() {
	fmt.Println(addStrings("17", "1222"))
}
