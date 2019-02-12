package main

import "fmt"
import "sort"

/*

Given array nums = [-1, 0, 1, 2, -1, -4],
Find all unique triplets which sum to 0

A solution set is:
[
  [-1, 0, 1],
  [-1, -1, 2]
]

Key Idea:
The naive approach of nested looping thrice and generating all possible triplets produces too many duplicate triplets.
So, we first sort the input which places equal numbers next to each other.
Then we have an outer loop index i going from 0 till len(arr)-2
For every outer loop index, we scan the rest of the array for matching zero sums.
Everytime we encounter duplicate element, either in the outer loop or in the inner scan, we skip elements using
either arr[i-1] == arr[i] or arr[i] == arr[i+1].
Along the way, we simply collect all the matching triplets.
*/

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var ans [][]int
	for i := 0; i < len(nums)-2; i++ {
		lo, hi := i+1, len(nums)-1
		if i == 0 || i > 0 && nums[i] != nums[i-1] {
			for lo < hi {
				if nums[lo]+nums[hi]+nums[i] == 0 {
					ans = append(ans, []int{nums[lo], nums[hi], nums[i]})
					for lo < hi && nums[lo] == nums[lo+1] {
						lo++
					}
					for lo < hi && nums[hi] == nums[hi-1] {
						hi--
					}
					lo++
					hi--
				} else if nums[lo]+nums[hi]+nums[i] < 0 {
					lo++
				} else {
					hi--
				}
			}
		}
	}
	return ans
}

func main() {
	input := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(input))
}
