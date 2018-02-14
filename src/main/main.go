package main

const debug = true

func main() {
	board := Board{}
	board.CreateEmptyBoard()
	board.PrintBoard()
	board.AddShips()
	board.PrintBoard()
}
