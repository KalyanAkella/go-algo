package main

import "fmt"

func main() {
	str := "a foo  bar   bazz"
	n := len(str)
	result := ""
	buff := ""
	for i := n - 1; i >= 0; i-- {
		if str[i] == ' ' {
			if len(buff) > 0 {
				result = fmt.Sprintf("%s%s", result, buff)
				buff = ""
			}
			result = fmt.Sprintf("%s%s", result, string(str[i]))
		} else {
			buff = fmt.Sprintf("%s%s", string(str[i]), buff)
		}
	}

	if len(buff) > 0 {
		result = fmt.Sprintf("%s%s", result, buff)
	}
	fmt.Printf("'%s'\n", result)
}
