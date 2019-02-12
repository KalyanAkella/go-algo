package main

import "fmt"

func pow(num int, exp uint) int {
	result := 1
	for exp > 0 {
		result *= num
		exp--
	}
	return result
}

func baseN(num int, base uint) int {
	if base < 2 {
		panic("Base cannot be less than 2")
	}
	result := 0
	for i := uint(0); num > 0; i++ {
		digit := num % int(base)
		result += pow(10, i) * digit
		num = num / int(base)
	}
	return result
}

func main() {
	fmt.Println(baseN(35, 3))
	fmt.Println(baseN(50, 3))
	fmt.Println(baseN(50, 2))
}
