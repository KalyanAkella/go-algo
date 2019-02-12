package main

import "fmt"

/*
You are given K eggs, and you have access to a building with N floors from 1 to N.

Each egg is identical in function, and if an egg breaks, you cannot drop it again.

You know that there exists a floor F with 0 <= F <= N such that any egg dropped at a floor higher than F will break, and any egg dropped at or below floor F will not break.

Each move, you may take an egg (if you have an unbroken one) and drop it from any floor X (with 1 <= X <= N).

Your goal is to know with certainty what the value of F is.

What is the minimum number of moves that you need to know with certainty what F is, regardless of the initial value of F?

Input: K = 1, N = 2
Output: 2
Explanation:
Drop the egg from floor 1.  If it breaks, we know with certainty that F = 0.
Otherwise, drop the egg from floor 2.  If it breaks, we know with certainty that F = 1.
If it didn't break, then we know with certainty F = 2.
Hence, we needed 2 moves in the worst case to know what F is with certainty.
Example 2:

Input: K = 2, N = 6
Output: 3
Example 3:

Input: K = 3, N = 14
Output: 4


Note:

1 <= K <= 100
1 <= N <= 10000
*/

func superEggDrop(K int, N int) int {
	return maxMoves(K, 1, N)
}

func maxMoves(eggs, begFlr, endFlr int) int {
	if begFlr > endFlr || eggs == 0 {
		fmt.Println(eggs, begFlr, endFlr, 0)
		return 0
	}
	midFlr := begFlr + (endFlr-begFlr)>>1
	result := 1 + max(
		maxMoves(eggs-1, begFlr, midFlr-1),
		maxMoves(eggs, midFlr+1, endFlr))
	fmt.Println(eggs, begFlr, endFlr, result)
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println("vim-go")
}
