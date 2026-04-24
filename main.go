package main

import (
	"fmt"
	"log"
)

func main() {

	board, err := NewBoard()

	if err != nil {
		log.Fatal(err)
	}

	err = board.ParseFen("8/8/8/4k3/8/8/8/3K4 b - - 0 1")

	if err != nil {
		log.Fatal(err)
	}

	if _, err := board.Check(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(board.String())

}
