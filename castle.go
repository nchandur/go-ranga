package main

type CastleBit int

const (
	WKCA CastleBit = 1
	WQCA CastleBit = 2
	BKCA CastleBit = 4
	BQCA CastleBit = 8
)

func (c *CastleBit) Valid() bool {
	return *c >= 0 && *c < 16
}

func (c *CastleBit) String() string {
	var res string

	if *c&WKCA != 0 {
		res += "K"
	}

	if *c&WQCA != 0 {
		res += "Q"
	}
	if *c&BKCA != 0 {
		res += "k"
	}
	if *c&BQCA != 0 {
		res += "q"
	}

	return res
}
