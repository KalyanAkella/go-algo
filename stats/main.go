package main

import "fmt"

func mean(nums []int) int {
	var avg float32 = 0.0
	size := len(nums)
	for _, num := range nums {
		avg += (float32(num) / float32(size))
	}
	return int(avg)
}

func std_dev(nums []int) int {
}

func mean_stream(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	var avg float32 = float32(nums[0])
	i := 1
	for {
		if i == len(nums) {
			return int(avg)
		}
		avg = (float32(nums[i]) + float32(nums[i-1])*avg) / float32(i+1)
		i++
	}
	return -1
}

func main() {
	fmt.Println(mean([]int{1, 2, 3, 4, 5}))
	fmt.Println(mean_stream([]int{1, 2, 3, 4, 5}))
}
