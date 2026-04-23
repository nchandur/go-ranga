package main

import "fmt"

type Piece int

const (
	Empty Piece = iota
	wP
	wN
	wB
	wR
	wQ
	wK
	bP
	bN
	bB
	bR
	bQ
	bK
	Offboard
)

func (p *Piece) Valid() bool {
	return *p >= wP && *p <= bK
}

func (p *Piece) ValidEmpty() bool {
	return p.Valid() || *p == Empty
}

func (p *Piece) String() string {
	if p.Valid() {
		return fmt.Sprintf("%c", PieceChar[*p])
	}
	return ""
}
