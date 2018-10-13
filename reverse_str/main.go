package main

import "fmt"

func reverse_recur(str string) string {
	if len(str) <= 1 {
		return str
	}
	return reverse_recur(str[1:]) + string(str[0])
}

func reverse(str string) string {
	ls := len(str)
	result := make([]rune, ls)
	for i := range str {
		result[ls-i-1] = rune(str[i])
	}
	return string(result)
}

func main() {
	fmt.Println(reverse_recur("abcd"))
	fmt.Println(reverse("abcd"))
}
