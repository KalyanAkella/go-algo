package main

import "fmt"

func print_table(n int) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			fmt.Printf("%d ", i*j)
		}
		fmt.Println()
	}
}

func main() {
	print_table(12)
}
