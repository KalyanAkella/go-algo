package main

import "fmt"

type Color byte

const (
	RED   Color = 1
	BLUE        = 2
	GREEN       = 3
)

func paintFill(screen [][]Color, row, col int, oldColor, newColor Color) {
	if row >= len(screen) || col >= len(screen[0]) {
		return
	}

	if screen[row][col] != oldColor || screen[row][col] == newColor {
		return
	}

	screen[row][col] = newColor
	offsets := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for _, offset := range offsets {
		paintFill(screen, row+offset[0], col+offset[1], oldColor, newColor)
	}
}

func printScreen(screen [][]Color) {
	for _, row := range screen {
		for _, col := range row {
			fmt.Printf("%v ", col)
		}
		fmt.Println()
	}
	fmt.Println()
}

func main() {
	screen := [][]Color{
		{RED, RED, RED, RED, RED},
		{RED, BLUE, RED, RED, RED},
		{RED, RED, RED, RED, RED},
		{RED, RED, BLUE, BLUE, RED},
		{RED, RED, RED, BLUE, RED},
	}
	printScreen(screen)
	paintFill(screen, 3, 2, BLUE, GREEN)
	printScreen(screen)
}
