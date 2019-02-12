package main

import "fmt"

func numWaysOld(numSteps uint) uint {
	switch numSteps {
	case 1:
		return 1
	case 2:
		return 2
	case 3:
		return 4
	default:
		return numWaysOld(numSteps-1) + numWaysOld(numSteps-2) + numWaysOld(numSteps-3)
	}
}

func numWaysBottomUp(numSteps uint) uint {
	memo := make([]uint, numSteps+1)
	memo[0] = 0
	memo[1] = 1
	memo[2] = 2
	memo[3] = 4
	for i := uint(4); i <= numSteps; i++ {
		memo[i] = memo[i-1] + memo[i-2] + memo[i-3]
	}
	return memo[numSteps]
}

func numWaysTopDown(numSteps uint) uint {
	memo := make([]uint, numSteps+1)
	return numWaysMemo(numSteps, memo)
}

func numWaysMemo(numSteps uint, memo []uint) uint {
	if memo[numSteps] == 0 {
		switch numSteps {
		case 1:
			memo[numSteps] = 1
		case 2:
			memo[numSteps] = 2
		case 3:
			memo[numSteps] = 4
		default:
			memo[numSteps] = numWaysMemo(numSteps-1, memo) + numWaysMemo(numSteps-2, memo) + numWaysMemo(numSteps-3, memo)
		}
	}
	return memo[numSteps]
}

func main() {
	for i := uint(4); i <= 100; i++ {
		fmt.Println("Steps:", i, "->", "TopDown", numWaysTopDown(i), "BottomUp", numWaysBottomUp(i))
	}
}
