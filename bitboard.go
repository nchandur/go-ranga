package main

import "math/bits"

type Bitboard uint64

// pops least significant bit from bitboard
func (b *Bitboard) PopBit() int {
	*b &= (*b - 1)
	return bits.TrailingZeros64(uint64(*b))
}

// counts number of bits set to 1 in bitboard
func (b *Bitboard) Count() int {
	return bits.OnesCount64(uint64(*b))
}
