package sorts

import (
	"fmt"
	"reflect"
	"testing"
)

func assert(t *testing.T, expected, actual interface{}) {
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Mismatch: Expected - %v, Actual - %v", expected, actual)
	}
}

func TestPartition(t *testing.T) {
	arr := []int{7, 2, 1, 6, 8, 5, 3, 4}
	pivot := partition(arr, 0, len(arr)-1)
	assert(t, 3, pivot)
	assert(t, []int{2, 1, 3, 4, 8, 5, 7, 6}, arr)
}

func TestInsertionSort(t *testing.T) {
	assert(t, []int{}, insertion_sort([]int{}))
	assert(t, []int{2, 7}, insertion_sort([]int{7, 2}))
	assert(t, []int{1, 2, 3, 5, 6, 7}, insertion_sort([]int{7, 2, 3, 1, 6, 5}))
	assert(t, []int{1, 2, 3, 4, 5, 6, 7}, insertion_sort([]int{7, 6, 5, 4, 3, 2, 1}))
}

func TestMerge(t *testing.T) {
	assert(t, []int{1, 2, 3}, merge([]int{}, []int{1, 2, 3}))
	assert(t, []int{1, 2, 3}, merge([]int{1}, []int{2, 3}))
	assert(t, []int{1, 2, 3}, merge([]int{1, 2}, []int{3}))
	assert(t, []int{1, 2, 3}, merge([]int{1, 2, 3}, []int{}))
}

func TestMergeSort(t *testing.T) {
	assert(t, []int{}, merge_sort([]int{}))
	assert(t, []int{2, 7}, merge_sort([]int{7, 2}))
	assert(t, []int{1, 2, 3, 5, 6, 7}, merge_sort([]int{7, 2, 3, 1, 6, 5}))
	assert(t, []int{1, 2, 3, 4, 5, 6, 7}, merge_sort([]int{7, 6, 5, 4, 3, 2, 1}))
}

func TestQuickSort(t *testing.T) {
	assert(t, []int{}, quick_sort([]int{}))
	assert(t, []int{2, 7}, quick_sort([]int{7, 2}))
	assert(t, []int{1, 2, 3, 5, 6, 7}, quick_sort([]int{7, 2, 3, 1, 6, 5}))
	assert(t, []int{1, 2, 3, 4, 5, 6, 7}, quick_sort([]int{7, 6, 5, 4, 3, 2, 1}))
}

func TestRandom(t *testing.T) {
	for i := 0; i < 1000; i++ {
		fmt.Printf("%d ", random(3, 10))
	}
	fmt.Println()
}
