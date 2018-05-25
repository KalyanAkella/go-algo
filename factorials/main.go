package main

import "fmt"

func factorial(num int) int {
	if num > 0 {
		return num * factorial(num-1)
	} else {
		return 1
	}
}

func main() {
	fmt.Println(factorial(4))
}
