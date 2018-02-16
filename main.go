package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

const debug = false

var reader *bufio.Reader

func main() {
	reader = bufio.NewReader(os.Stdin)
	board := NewBoard()
	board.AddShips()
	remainingShips := board.GetRemainingShips()

	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()

	for len(remainingShips) > 0 {
		fmt.Printf("At least one ship exists for each of the following sizes: %v\n", remainingShips)
		board.PrintBoard()

		row, col := getTargetLocation()

		if row == -1 && col == -1 {
			break
		}

		c := exec.Command("clear")
		c.Stdout = os.Stdout
		c.Run()

		switch board.GetValueAtLoc(row, col) {
		case ' ':
			fmt.Println("You missed!")
			board.SetValue(row, col, 'O')
		case 'X', 'O':
			fmt.Println("Already guessed this location...")
		default:
			fmt.Println("HIT!!!")
			board.SetValue(row, col, 'X')
		}

		remainingShips = board.GetRemainingShips()
	}

	if len(board.GetRemainingShips()) == 0 {
		fmt.Println("You win!")
	} else {
		fmt.Println("Exiting game...")
	}
}

func getTargetLocation() (row, col int) {
	text := ""

	for len(text) < 2 {
		fmt.Print("Enter coordinates (i.e. A1) or \"QUIT\" to exit: ")
		text, _ = reader.ReadString('\n')
		text = text[:len(text)-1]
	}

	if text == "QUIT" {
		return -1, -1
	}

	row = int(strings.ToUpper(text)[0] - 'A')
	col, err := strconv.Atoi(text[1:])

	if row < 0 || row >= 10 {
		panic(fmt.Sprintf("Letter \"%c\" is out of range!", 'A'+row))
		// goto RETRY
	}

	if err != nil {
		panic(fmt.Sprintf("Second value (%s) is not a number!", text[1:2]))
		// goto RETRY
	} else if col = col - 1; col < 0 || col >= 10 {
		panic(fmt.Sprintf("Second value (%s) must be between [0, 9]!", text[1:2]))
		// goto RETRY
	}

	return row, col
}
