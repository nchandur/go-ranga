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

func NewBoard() Board {
	board := Board{}

	for idx := range board.KingsSquare {
		board.KingsSquare[idx] = NoSquare
	}

	board.EnPassant = NoSquare

	return board

}

func (b *Board) Reset() error {
	for idx := range BOARD_SQ_NUM {
		b.Pieces[idx] = Offboard
	}

	for idx := range 64 {
		sq64, err := SQ120(Square(idx))

		if err != nil {
			return fmt.Errorf("failed to reset board: %v", err)
		}

		b.Pieces[sq64] = Empty
	}

	b.BigPiece = [2]int{}
	b.MajorPiece = [2]int{}
	b.MinorPiece = [2]int{}
	b.Material = [2]int{}

	b.Pawns = [3]Bitboard{}
	b.PieceNum = [13]int{}

	b.KingsSquare[White] = NoSquare
	b.KingsSquare[Black] = NoSquare

	b.Side = Both
	b.EnPassant = NoSquare

	b.Ply = 0
	b.HistoryPly = 0
	b.CastleBit = 0

	b.PositionKey = uint64(0)
	b.History = [MAX_GAMES_MOVES]History{}

	return nil
}

func (b *Board) GeneratePositionKey() (uint64, error) {

	var res uint64

	for sq := range BOARD_SQ_NUM {
		piece := b.Pieces[sq]

		if piece != Empty {
			if !piece.Valid() {
				return uint64(0), fmt.Errorf("failed to generate position key. invalid piece: %d", piece)
			}
			res ^= PieceKeys[piece][sq]
		}
	}

	if b.Side == White {
		res ^= SideKey
	}

	if b.EnPassant != NoSquare {
		if !b.EnPassant.Valid() {
			return uint64(0), fmt.Errorf("failed to generate position key. invalid en passant square %d", b.EnPassant)
		}

		if !b.EnPassant.IsOnBoard() {
			return uint64(0), fmt.Errorf("failed to generate position key. en passant square %d is offboard", b.EnPassant)
		}

		if FileBoard[b.EnPassant] != 3 && FileBoard[b.EnPassant] != 6 {
			return uint64(0), fmt.Errorf("failed to generate position key. en passant square must be on Rank 3 or 6, found on %d", FileBoard[b.EnPassant])
		}

		res ^= PieceKeys[Empty][b.EnPassant]

	}

	if !b.CastleBit.Valid() {
		return uint64(0), fmt.Errorf("failed to generate position key. invalid castle bit: %d", b.CastleBit)
	}

	res ^= CastleKey[b.CastleBit]

	return res, nil
}

func (b *Board) String() string {
	var builder strings.Builder

	fmt.Fprintf(&builder, "\n====Game Board====\n")

	for rank := Rank8; rank >= Rank1; rank-- {
		fmt.Fprintf(&builder, "%d  ", rank+1)

		for file := FileA; file <= FileH; file++ {
			sq, _ := FRTo120(file, rank)
			piece := b.Pieces[sq]

			fmt.Fprintf(&builder, "%3s ", &piece)
		}
		fmt.Fprintf(&builder, "\n")
	}

	fmt.Fprintf(&builder, "\n   ")

	for file := FileA; file <= FileH; file++ {
		fmt.Fprintf(&builder, "%3c", 'a'+file)
	}

	fmt.Fprintf(&builder, "\n")

	fmt.Fprintf(&builder, "Side: %s", &b.Side)
	fmt.Fprintf(&builder, "En Passant: %s", &b.EnPassant)
	fmt.Fprintf(&builder, "Castle: %s", &b.CastleBit)
	fmt.Fprintf(&builder, "Position Key: %x\n", b.PositionKey)

	return builder.String()
}
