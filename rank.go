package main

type Rank int

const (
	Rank1 Rank = iota
	Rank2
	Rank3
	Rank4
	Rank5
	Rank6
	Rank7
	Rank8
	RankNone
)

func (r *Rank) Valid() bool {
	return *r >= Rank1 && *r <= Rank8

}
