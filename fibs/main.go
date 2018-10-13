package main

import "fmt"

func nth_fib(n int) uint64 {
	if n < 2 {
		return uint64(n)
	}
	p, c := uint64(0), uint64(1)
	for n >= 2 {
		t := uint64(c)
		c = p + c
		p = t
		n = n - 1
	}
	return c
}

func nth_fib_recur(n int) uint64 {
	if n < 2 {
		return uint64(n)
	} else {
		return nth_fib_recur(n-1) + nth_fib_recur(n-2)
	}
}

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d,", nth_fib(i))
	}
	fmt.Println("\n======================================================")
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d,", nth_fib_recur(i))
	}
	fmt.Println()
}
