package main

import (
	"fmt"
	"math/rand"
)

/*
arr = [1,2,3,4,5,6,7,8]
m = 5
*/

func randomSet(arr []int, m int) []int {
	var ans []int
	for i := len(arr) - 1; m > 0; i, m = i-1, m-1 {
		randIndex := rand.Intn(i)
		ans = append(ans, arr[randIndex])
		arr[randIndex], arr[i] = arr[i], arr[randIndex]
	}
	return ans
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8}
	m := 5
	for i := 1; i <= 10; i++ {
		fmt.Println(randomSet(arr, m))
	}
}
