package main

import (
	"log"
	"math/rand/v2"
)

func initSquareIndices() {
	sq64 := Square(0)

	for idx := range BOARD_SQ_NUM {
		Sq120To64[idx] = 65
	}

	for idx := range 64 {
		Sq64To120[idx] = NoSquare
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

func initBitMasks() {

	for idx := range 64 {
		SetMask[idx] = uint64(0)
		ClearMask[idx] = uint64(0)
	}

	for idx := range 64 {
		SetMask[idx] = uint64(1) << idx
		ClearMask[idx] = ^SetMask[idx]
	}

}

func initHashkeys() {

	for idx := range 13 {
		for jdx := range 120 {
			PieceKeys[idx][jdx] = rand.Uint64()
		}
	}

	SideKey = rand.Uint64()

	for idx := range 16 {
		CastleKey[idx] = rand.Uint64()
	}

}

func initFileRankBoard() {

	for idx := range BOARD_SQ_NUM {
		FileBoard[idx] = FileNone
		RankBoard[idx] = RankNone
	}

	for rank := Rank1; rank <= Rank8; rank++ {
		for file := FileA; file <= FileH; file++ {
			sq, err := FRTo120(file, rank)

			if err != nil {
				log.Fatal(err)
			}

			FileBoard[sq] = file
			RankBoard[sq] = rank
		}
	}
}

func init() {
	initSquareIndices()
	initBitMasks()
	initHashkeys()
	initFileRankBoard()
}
