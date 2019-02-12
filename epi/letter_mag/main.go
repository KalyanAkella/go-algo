package main

import "fmt"

func isPossible(letter, magazine string) bool {
	freqs := make(map[byte]int)

	for i := range letter {
		c := letter[i]
		if n, present := freqs[c]; present {
			freqs[c] = n + 1
		} else {
			freqs[c] = 1
		}
	}

	for i := range magazine {
		c := magazine[i]
		if n, present := freqs[c]; present {
			freqs[c] = n - 1
			if freqs[c] == 0 {
				delete(freqs, c)
				if len(freqs) == 0 {
					break
				}
			}
		}
	}
	return len(freqs) == 0
}

func main() {
	magazine := "a quick brown fox jumped over the lazy dog"
	fmt.Println(isPossible("a dog", magazine))
	fmt.Println(isPossible("c program", magazine))
	fmt.Println(isPossible("hello world", magazine))
}
