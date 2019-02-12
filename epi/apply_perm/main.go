package main

import "fmt"

// In place permutation application
func applyPerm(arr, perm []int) {
	for i := range perm {
		if perm[i] >= 0 {
			a, temp := i, arr[i]
			for {
				next_a, next_temp := perm[a], arr[perm[a]]
				arr[next_a] = temp
				perm[a] -= len(perm)
				a, temp = next_a, next_temp
				if a == i {
					break
				}
			}
		}
	}
}

func main() {
	arr := []int{10, 5, 11, 7}
	perm := []int{2, 0, 1, 3}
	fmt.Println(arr)
	applyPerm(arr, perm)
	fmt.Println(arr)
}
