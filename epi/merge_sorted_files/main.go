package main

import (
	"bufio"
	"container/heap"
	"io"
	"log"
	"os"
	"strconv"
)

type IntHeap []int

func (this IntHeap) Len() int {
	return len(this)
}

func (this IntHeap) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func (this IntHeap) Less(i, j int) bool {
	return this[i] < this[j]
}

func (this *IntHeap) Push(x interface{}) {
	num := x.(int)
	*this = append(*this, num)
}

func (this *IntHeap) Pop() interface{} {
	old := *this
	n := len(old)
	top := old[n-1]
	*this = old[0 : n-1]
	return top
}

func insertAllLinesIntoHeap(h *IntHeap, fd *os.File) error {
	reader := bufio.NewScanner(bufio.NewReader(fd))
	reader.Split(bufio.ScanLines)
	for reader.Scan() {
		line := reader.Text()
		num, err := strconv.Atoi(line)
		if err != nil {
			return err
		}
		heap.Push(h, num)
	}
	return nil
}

func writeAllLinesFromHeap(h *IntHeap, fd *os.File) error {
	if _, err := fd.Seek(0, 0); err != nil {
		return err
	}
	writer := bufio.NewWriter(fd)
	for (*h).Len() > 0 {
		num := heap.Pop(h).(int)
		numStr := strconv.Itoa(num) + "\n"
		if _, err := writer.WriteString(numStr); err != nil {
			return err
		}
		writer.Flush()
	}
	return nil
}

func mergeSortedFiles(resultFile, inputFile string) error {
	resultFD, err := os.OpenFile(resultFile, os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		return err
	}
	defer resultFD.Close()
	inputFD, err := os.OpenFile(inputFile, os.O_RDONLY, 0)
	if err != nil {
		return err
	}
	defer inputFD.Close()

	h := &IntHeap{}
	heap.Init(h)

	if err := insertAllLinesIntoHeap(h, resultFD); err != nil && err != io.EOF {
		return err
	}
	if err := insertAllLinesIntoHeap(h, inputFD); err != nil && err != io.EOF {
		return err
	}
	if err := writeAllLinesFromHeap(h, resultFD); err != nil && err != io.EOF {
		return err
	}
	return nil
}

func main() {
	resultFile := "/tmp/merge/result.txt"
	files := []string{"/tmp/merge/file1.txt", "/tmp/merge/file2.txt"}
	for _, file := range files {
		if err := mergeSortedFiles(resultFile, file); err != nil {
			log.Fatal(err)
		}
	}
}
