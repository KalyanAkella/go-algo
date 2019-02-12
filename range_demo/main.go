package main

import "fmt"

func main() {
	arr := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
	}

	for _, nums := range arr {
		fmt.Println(nums)
	}

	fmt.Println("vim-go")
}
