package main

import "fmt"

func pow(num byte, exp uint) uint64 {
	result := uint64(1)
	for ; exp > 0; exp-- {
		result *= uint64(num)
	}
	return result
}

// Doesn't support negative numbers
func convertToDecimal(numStr string, b1 byte) uint64 {
	result := uint64(0)
	for i := len(numStr) - 1; i >= 0; i-- {
		digitByte := numStr[i]
		var digit byte
		if digitByte >= 'A' && digitByte <= 'F' {
			digit = 10 + digitByte - 'A'
		} else {
			digit = digitByte - '0'
		}
		result = result + uint64(digit)*pow(b1, uint(len(numStr)-i-1))
	}
	return result
}

func convertFromDecimal(input uint64, b2 byte) string {
	result := ""
	for input > 0 {
		digit := input % uint64(b2)
		var digitStr string
		if digit > 9 {
			digitStr = string('A' + digit - 10)
		} else {
			digitStr = string('0' + digit)
		}
		result = digitStr + result
		input = input / uint64(b2)
	}
	return result
}

func convert(str1 string, b1, b2 byte) string {
	input := convertToDecimal(str1, b1)
	return convertFromDecimal(input, b2)
}

func main() {
	fmt.Println(convert("F", 16, 2))
	fmt.Println(convert("1EF", 16, 2))
	fmt.Println(convert("111101111", 2, 16))
}
