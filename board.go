package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

var random rand.Rand

const boardLetters = "ABCDEFGHIJ"

var ships = []int{5, 4, 3, 3, 2}

type Board struct {
	grid [10][10]byte
}

func NewBoard() *Board {
	s1 := rand.NewSource(time.Now().UnixNano())
	random = *(rand.New(s1))

	board := new(Board)
	// Empty board
	for rowIndex, row := range board.grid {
		for cell := range row {
			board.grid[rowIndex][cell] = ' '
		}
	}

	return board
}

func (board *Board) AddShips() {
	for _, ship := range ships {
		isValid := false
		var row, col, direction int
		for !isValid {
			row, col, direction = generateRandomShipPositions()
			// fmt.Printf("%c%d\n", boardLetters[row], col)
			isValid = board.CheckIfShipPositionIsValid(row, col, direction, ship)
			if debug {
				fmt.Printf("%d is valid\n", ship)
			}
		}
		// valid position. ship can be set now
		board.AddShip(row, col, direction, ship)
		if debug {
			fmt.Printf("Ship size %d added at %c%d!\n", ship, boardLetters[row], col)
		}
	}
}

func (board *Board) SetValue(row, col int, value byte) bool {
	if currentVal := board.GetValueAtLoc(row, col); currentVal == 'X' || currentVal == 'O' {
		return false
	}
	board.grid[row][col] = value
	return true
}

func (board *Board) AddShip(row, col, direction, shipSize int) {
	shipChar := strconv.Itoa(shipSize)[0]
	switch direction {
	case 0:
		for startRow := row; startRow >= 0 && startRow > row-shipSize; startRow-- {
			if !board.SetValue(startRow, col, shipChar) && debug {
				fmt.Println("OVERWRITING VALUE")
			}
		}
	case 1:
		for startCol := col; startCol < 10 && startCol < col+shipSize; startCol++ {
			if !board.SetValue(row, startCol, shipChar) && debug {
				fmt.Println("OVERWRITING VALUE")
			}
		}
	case 2:
		for startRow := row; startRow < 10 && startRow < row+shipSize; startRow++ {
			if !board.SetValue(startRow, col, shipChar) && debug {
				fmt.Println("OVERWRITING VALUE")
			}
		}
	case 3:
		for startCol := col; startCol >= 0 && startCol > col-shipSize; startCol-- {
			if !board.SetValue(row, startCol, shipChar) && debug {
				fmt.Println("OVERWRITING VALUE")
			}
		}
	}
}

func (board *Board) PrintBoard() {
	fmt.Println("     1   2   3   4   5   6   7   8   9  10")
	fmt.Println("   -----------------------------------------")

	for rowIndex, row := range board.grid {
		fmt.Printf(" %c |", boardLetters[rowIndex])
		for _, cell := range row {
			if debug {
				fmt.Printf(" %c |", cell)
			} else {
				switch cell {
				case 'O', 'X':
					fmt.Printf(" %c |", cell)
				default:
					fmt.Print("   |")
				}
			}
		}
		fmt.Println("\n   -----------------------------------------")
	}
}

func (board *Board) GetValueAtLoc(row, col int) (value byte) {
	return board.grid[row][col]
}

func (board *Board) IsEmptyLocation(row, col int) (isEmpty bool) {
	return board.grid[row][col] == ' '
}

func generateRandomShipPositions() (row, col, direction int) {
	row = random.Intn(10)
	col = random.Intn(10)
	direction = random.Intn(4)
	return
}

func (board *Board) CheckIfShipPositionIsValid(row, col, direction, shipSize int) (isValid bool) {
	switch direction {
	case 0:
		// up
		if row-shipSize < 0 {
			return false
		}
		// if there is already some ship in any of the proposed spaces, return false
		for startRow := row; startRow >= 0 && startRow > row-shipSize; startRow-- {
			if !board.IsEmptyLocation(startRow, col) {
				return false
			}
		}

	case 1:
		// right
		if col+shipSize > 10 {
			return false
		}

		for startCol := col; startCol < 10 && startCol < col+shipSize; startCol++ {
			if !board.IsEmptyLocation(row, startCol) {
				return false
			}
		}

	case 2:
		// down
		if row+shipSize > 10 {
			return false
		}

		for startRow := row; startRow < 10 && startRow < row+shipSize; startRow++ {
			if !board.IsEmptyLocation(startRow, col) {
				return false
			}
		}

	case 3:
		// left
		if col-shipSize < 0 {
			return false
		}

		for startCol := col; startCol >= 0 && startCol > col-shipSize; startCol-- {
			if !board.IsEmptyLocation(row, startCol) {
				return false
			}
		}
	}
	return true
}

func isIntInList(value int, list []int) bool {
	for _, num := range list {
		if num == value {
			return true
		}
	}
	return false
}

func getUniqueShipSizes(shipArray []int) []int {
	var uniqueShips []int
	for _, shipSize := range ships {
		if !isIntInList(shipSize, uniqueShips) {
			uniqueShips = append(uniqueShips, shipSize)
		}
	}
	return uniqueShips
}

func (board *Board) GetRemainingShips() (remainingShips []int) {
	uniqueShips := getUniqueShipSizes(ships)
	alive := make([]bool, len(uniqueShips))
	for _, row := range board.grid {
	CELL_LOOP:
		for _, cell := range row {
			for shipIndex, ship := range uniqueShips {
				if strconv.Itoa(ship)[0] == cell {
					alive[shipIndex] = true
					continue CELL_LOOP
				}
			}
		}
	}

	for index, isAlive := range alive {
		if isAlive {
			remainingShips = append(remainingShips, uniqueShips[index])
		}
	}
	return
}
