package main

func gen(arr []int) [][]int {
	la := len(arr)
	k := la - 2
	for k >= 0 && arr[k] > arr[k+1] {
		k--
	}
	if k < 0 {
		return nil
	}
	return nil
}

func generate(num int) [][]int {
	arr := make([]int, num)
	for i := 0; i < num; i++ {
		arr[i] = i + 1
	}
	return gen(arr)
}

func main() {
	generate(5)
}
