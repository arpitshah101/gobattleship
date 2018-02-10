package main

import (
	"fmt"
)

func main() {
	board := createEmptyBoard()
	printBoard(board)
}

func createEmptyBoard() (board [10][10]byte) {
	// Empty board
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			board[i][j] = ' '
		}
	}
	return
}

func printBoard(board [10][10]byte) {
	fmt.Println("    1   2   3   4   5   6   7   8   9  10")
	fmt.Println("  -----------------------------------------")

	letters := "ABCDEFGHIJ"

	for i := 0; i < 10; i++ {
		fmt.Printf("%c |", letters[i])
		for j := 0; j < 10; j++ {
			fmt.Printf(" %c |", board[i][j])
		}
		fmt.Println("\n  -----------------------------------------")
	}
}
