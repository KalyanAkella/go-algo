package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

/*
input is a file having 4 billion non-negative integers
2^32 -> 2^31 non-negative numbers
2^31 -> ~2 billion, but input file has 4 billion -> there are duplicates

1GB memory available -> 1 billion bytes -> 8 billion bits
0.5GB -> 4 billion bits

Approach:
1. allocate 4 billion bits in a bit field
2. scan through the file and mark the bit position of the current number in the bit field
3. at the end, scan through the bit field and return the first bit position that is unset
*/

type BitField struct {
	bytes []byte
}

func createBitField(sizeInBytes uint) *BitField {
	return &BitField{bytes: make([]byte, sizeInBytes)}
}

func openFile(file string) *os.File {
	if file, err := os.Open(file); err != nil {
		panic(err)
	} else {
		return file
	}
}

func markInt(num int, bitField *BitField) {
	byteOffset := num >> 3 // divide by 8
	bitOffsetWithinByte := uint(num % 8)
	bitField.bytes[byteOffset] |= (1 << bitOffsetWithinByte) // TODO: Bounds check ? (also LSB order)
}

func firstUnSetBit(bitField *BitField) int {
	for i, bits := range bitField.bytes {
		for j := uint(0); j < 8; j++ {
			temp := bits & (1 << j)
			if temp == 0 {
				return i*8 + int(j)
			}
		}
	}
	return -1
}

func scanAndMarkInts(file *os.File, bitField *BitField) {
	fileScan := bufio.NewScanner(bufio.NewReader(file))
	fileScan.Split(bufio.ScanLines)
	for fileScan.Scan() {
		numStr := fileScan.Text()
		if num, err := strconv.Atoi(numStr); err != nil {
			panic(err) // TODO: Should we just skip ?
		} else {
			markInt(num, bitField)
		}
	}
	if err := fileScan.Err(); err != nil {
		panic(err)
	}
}

func missingInt(file string) int {
	sizeInBytes := math.MaxInt32 >> 3
	bitField := createBitField(uint(sizeInBytes))
	fd := openFile(file)
	defer fd.Close()
	scanAndMarkInts(fd, bitField)
	return firstUnSetBit(bitField)
}

//TODO: Test this
func main() {
	fmt.Println("vim-go")
}
