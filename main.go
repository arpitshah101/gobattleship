package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const debug = false

var reader *bufio.Reader

func main() {
	reader = bufio.NewReader(os.Stdin)
	board := Board{}
	board.CreateEmptyBoard()
	board.PrintBoard()
	board.AddShips()
	board.PrintBoard()
}

func getTargetLocation() (row, col int) {
	fmt.Print("Enter coordinates (i.e. A1): ")
	text, _ := reader.ReadString('\n')
	text = text[:len(text)-1]

	row = int(strings.ToUpper(text)[0] - 'A')
	col, err := strconv.Atoi(text[1:2])

	if row < 0 || row >= 10 {
		panic(fmt.Sprintf("Letter \"%c\" is out of range!", 'A'+row))
	}

	if err != nil {
		panic(fmt.Sprintf("Second value (%s) is not a number!", text[1:2]))
	} else if col < 0 || col >= 10 {
		panic(fmt.Sprintf("Second value (%s) must be between [0, 9]!", text[1:2]))
	}

	return row, col
}
