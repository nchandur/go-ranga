package main

import (
	"fmt"
	"log"
)

func main() {

	board := NewBoard()
	err := board.ParseFen(START)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(&board)

}
