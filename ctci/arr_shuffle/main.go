package main

import (
	"fmt"
	"math/rand"
)

func shuffle(arr []int) {
	for i := len(arr) - 1; i >= 1; i-- {
		randIndex := rand.Intn(i)
		arr[randIndex], arr[i] = arr[i], arr[randIndex]
	}
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	fmt.Println(arr)
	shuffle(arr)
	fmt.Println(arr)
}
