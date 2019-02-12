package main

import (
	"bufio"
	"fmt"
	"sort"
	"strings"
)

func sortLines(scanner *bufio.Scanner) []string {
	linesByLength := make(map[int][]string)
	for scanner.Scan() {
		line := scanner.Text()
		lineLen := len(line)
		if lines, present := linesByLength[lineLen]; present {
			linesByLength[lineLen] = append(lines, line)
		} else {
			linesByLength[lineLen] = []string{line}
		}
	}

	var lens []int
	for length, lines := range linesByLength {
		lens = append(lens, length)
		sort.Strings(lines)
	}
	sort.Ints(lens)

	var result []string
	for _, l := range lens {
		result = append(result, linesByLength[l]...)
	}
	return result
}

func main() {
	lines :=
		`abcd
xyz
pq
ab
abcdef
abc
def`
	scanner := bufio.NewScanner(strings.NewReader(lines))
	scanner.Split(bufio.ScanLines)

	fmt.Println(sortLines(scanner))
}
