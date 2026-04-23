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
