package main

import "fmt"

func isMultipleOf3(str string) bool {
	s := 0
	for i := range str {
		num := int(str[i] - '0')
		s = s*2 + num
	}
	return s%3 == 0
}

func main() {
	fmt.Println(isMultipleOf3("011"))
	fmt.Println(isMultipleOf3("100"))
}
