package main

import (
	"fmt"
	"math/rand"
	"time"
)

func random(floor, ceiling int) int {
	return floor + rand.Intn(ceiling+1-floor)
}

func shuffle(arr []int) {
	for i := 0; i < len(arr); i++ {
		rand_idx := random(i, len(arr)-1)
		arr[i], arr[rand_idx] = arr[rand_idx], arr[i]
	}
}

func test_shuffle(arr []int) {
	fmt.Println("Before shuffling:", arr)
	shuffle(arr)
	fmt.Println("After shuffling:", arr)
}

func main() {
	rand.Seed(time.Now().Unix())
	test_shuffle([]int{4, 5, 6, 7, 8, 9})
	test_shuffle([]int{4, 5})
	test_shuffle([]int{4})
}
