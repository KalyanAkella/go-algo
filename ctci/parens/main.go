package main

import "fmt"

type Empty struct{}
type StringSet map[string]Empty

func generateParens(num int) StringSet {
	ans := make(StringSet)
	if num == 0 {
		ans[""] = Empty{}
		return ans
	}
	subParens := generateParens(num - 1)
	for subParen := range subParens {
		for i := 0; i <= len(subParen); i++ {
			paren := subParen[0:i] + "()" + subParen[i:]
			ans[paren] = Empty{}
		}
	}
	return ans
}

func main() {
	fmt.Println(generateParens(1))
	fmt.Println(generateParens(2))
	fmt.Println(generateParens(3))
	fmt.Println(generateParens(4))
	fmt.Println(generateParens(10))
}
