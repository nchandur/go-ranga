package main

import (
	"fmt"
	"strings"
)

type Board struct {
	Pieces [BOARD_SQ_NUM]Piece
	Pawns  [3]Bitboard

	KingsSquare [2]Square

	Side      Color
	EnPassant Square

	FiftyMove  int
	Ply        int
	HistoryPly int

	PositionKey uint64

	CastleBit
	PieceNum   [13]int
	BigPiece   [2]int
	MajorPiece [2]int
	MinorPiece [2]int
	Material   [2]int

	History [MAX_GAMES_MOVES]History

	// 13: piece number each side
	// 10: squares on which pieces exist
	PieceList [13][10]int
}

type History struct {
	Move int
	CastleBit
	EnPassant   Square
	FiftyMove   int
	PositionKey uint64
}

func NewBoard() (*Board, error) {
	board := Board{}

	if err := board.Reset(); err != nil {
		return nil, fmt.Errorf("failed to generate new board: %v", err)
	}

	return &board, nil

}

func (b *Board) String() string {
	var builder strings.Builder

	builder.WriteString("\n  +---------------------------------+\n")

	for rank := Rank8; rank >= Rank1; rank-- {
		fmt.Fprintf(&builder, "%d | ", rank+1)

		for file := FileA; file <= FileH; file++ {
			sq, _ := FRTo120(file, rank)
			piece := b.Pieces[sq]

			if piece == Empty {
				builder.WriteString(" .  ")
			} else {
				fmt.Fprintf(&builder, " %s  ", &piece)
			}
		}
		builder.WriteString("|\n")
	}

	builder.WriteString("  +---------------------------------+\n")

	builder.WriteString("     a   b   c   d   e   f   g   h\n\n")

	fmt.Fprintf(&builder, "Side To Move : %s\n", &b.Side)
	fmt.Fprintf(&builder, "En Passant   : %s\n", &b.EnPassant)
	fmt.Fprintf(&builder, "Castling     : %s\n", &b.CastleBit)
	fmt.Fprintf(&builder, "Pos Key      : %x\n", b.PositionKey)

	return builder.String()
}
