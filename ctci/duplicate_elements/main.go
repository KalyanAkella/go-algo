package main

import (
	"fmt"
	"math/rand"
)

type BitField struct {
	bytes []byte
}

func createBitField(sizeInBytes uint) *BitField {
	return &BitField{bytes: make([]byte, sizeInBytes)}
}

func (bf *BitField) markNumbers(arr []int, printDups bool) {
	for _, num := range arr {
		bitfieldOffset := num / 8
		bitoffsetInByte := uint(num % 8)
		if bf.bytes[bitfieldOffset]&(1<<bitoffsetInByte) == 0 {
			bf.bytes[bitfieldOffset] |= (1 << bitoffsetInByte)
		} else {
			if printDups {
				fmt.Println(num)
			}
		}
	}
}

const AvailableMemInBytes = 4000

func printDups(arr []int, N int) {
	if N > AvailableMemInBytes*8 {
		panic("Not enough memory")
	}
	bitfield := createBitField(AvailableMemInBytes)
	bitfield.markNumbers(arr, true)
}

func main() {
	var arr []int
	N := 500
	for i := 0; i < 50; i++ {
		n := rand.Intn(N)
		arr = append(arr, n)
	}
	fmt.Println(arr)
	printDups(arr, N)
}
