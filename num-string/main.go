package main

import (
	"flag"
	"fmt"
	"math"
)

/*
Converts a num to string comma separated
1234567 -> 1,234,567
123456 -> 123,456
12345 -> 12,345
1234 -> 1,234
*/

var num uint

func init() {
	flag.UintVar(&num, "num", num, "the number to format")
	flag.Parse()
}

func format_US(num uint, places uint) string {
	result := ""
	var factor uint = uint(math.Pow10(int(places)))
	for num > 0 {
		val := num % factor
		result = fmt.Sprintf("%d,%s", val, result)
		num = num / factor
	}
	return result[:len(result)-1]
}

func main() {
	fmt.Println("Given number:", num)
	fmt.Println("Formatted number:", format_US(num, 3))
}
