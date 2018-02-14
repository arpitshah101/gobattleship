package main

import (
	"math/rand"
	"time"
)

var random rand.Rand
var debug bool

func main() {
	debug = true
	s1 := rand.NewSource(time.Now().UnixNano())
	random = *(rand.New(s1))
	board := Board{}
	board.CreateEmptyBoard()
	board.PrintBoard()
	board.AddShips()
	board.PrintBoard()
}
