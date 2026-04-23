package main

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
)

func (p *Piece) Valid() bool {
	return *p >= wP && *p <= bK
}

func (p *Piece) ValidEmpty() bool {
	return p.Valid() || *p == Empty
}
