package main

import "fmt"

func localMinima(arr []int) int {
	result := arr[0]
	for i := 1; i < len(arr); i++ {
		if i < len(arr)-1 {
			if arr[i-1] > arr[i] && arr[i] < arr[i+1] {
				result = arr[i]
				break
			}
		} else {
			if arr[i-1] > arr[i] {
				result = arr[i]
				break
			}
		}
	}
	return result
}

func main() {
	fmt.Println(localMinima([]int{11, 4, 2, 5, 11, 13, 5}))
	fmt.Println(localMinima([]int{1, 2, 3, 4}))
	fmt.Println(localMinima([]int{3}))
	fmt.Println(localMinima([]int{6, 4}))
}
