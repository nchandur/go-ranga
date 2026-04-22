package main

import "log"

func initSquareIndices() {
	sq64 := Square(0)

	for idx := range BOARD_SQ_NUM {
		Sq120To64[idx] = 65
	}

	for idx := range 64 {
		Sq64To120[idx] = Offboard
	}

	for rank := Rank1; rank <= Rank8; rank++ {
		for file := FileA; file <= FileH; file++ {
			sq, err := FRTo120(file, rank)

			if err != nil {
				log.Fatal(err)
			}

			Sq64To120[sq64] = sq
			Sq120To64[sq] = sq64
			sq64++
		}
	}

}

func init() {
	initSquareIndices()
}
