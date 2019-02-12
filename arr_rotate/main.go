package main

import "fmt"

func leftRotate(arr []int) {
	la := len(arr)
	first := arr[0]
	for i := 1; i < la; i++ {
		arr[i-1] = arr[i]
	}
	arr[la-1] = first
}

func rightRotate(arr []int) {
	la := len(arr)
	last := arr[la-1]
	for i := la - 2; i >= 0; i-- {
		arr[i+1] = arr[i]
	}
	arr[0] = last
}

func main() {
	arr := []int{1, 3, 5, 7, 9, 10}
	fmt.Println(arr)
	for range arr {
		leftRotate(arr)
		fmt.Println(arr)
	}
	for range arr {
		rightRotate(arr)
		fmt.Println(arr)
	}
}
