package main

import "fmt"

/*
Rearranges all elements in the input array 'arr' and given index, such that elements less than
arr[index] are present at the left, elements equal to arr[index] are in the middle and elements greter than
arr[index] are present at the right end of the array. Time complexity: O(n), space complexity: O(1)
*/
func rearrange(arr []int, pivot int) {
	smaller, equal, larger := 0, 0, len(arr)-1
	/*
		bottom = arr[0:smaller-1]
		middle = arr[smaller:equal-1]
		unclassified = arr[equal:larger]
		top = arr[larger+1:len(arr)-1]
	*/
	for equal <= larger {
		if arr[equal] < pivot {
			arr[equal], arr[smaller] = arr[smaller], arr[equal]
			smaller++
			equal++
		} else if arr[equal] > pivot {
			arr[equal], arr[larger] = arr[larger], arr[equal]
			larger--
		} else {
			equal++
		}
	}
}

func main() {
	arr := []int{2, 1, 4, 9, 6, 8, 5, 7, 3, 6}
	fmt.Println(arr)
	rearrange(arr, 6)
	fmt.Println(arr)
}
