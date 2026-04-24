package main

import (
	"fmt"
	"math/bits"
)

type Bitboard uint64

// pops least significant bit from bitboard
func (b *Bitboard) PopBit() int {
	res := bits.TrailingZeros64(uint64(*b))
	*b &= (*b - 1)
	return res
}

// counts number of bits set to 1 in bitboard
func (b *Bitboard) Count() int {
	return bits.OnesCount64(uint64(*b))
}

// set bit in bitboard
func (b *Bitboard) SetBit(square Square) error {
	if !square.Valid() {
		return fmt.Errorf("failed to set bit: invalid square. must be between 0 & 119")
	}

	sq64, err := SQ64(square)

	if err != nil {
		return fmt.Errorf("failed to set bit: %v", err)
	}

	*b |= Bitboard(SetMask[sq64])
	return nil
}

// clear bit from bitboard
func (b *Bitboard) ClearBit(square Square) error {
	if !square.Valid() {
		return fmt.Errorf("failed to set bit: invalid square. must be between 0 & 119")
	}

	sq64, err := SQ64(square)

	if err != nil {
		return fmt.Errorf("failed to set bit: %v", err)
	}

	*b &= Bitboard(ClearMask[sq64])
	return nil
}
