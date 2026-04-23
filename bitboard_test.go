package main

import "testing"

func TestBitboardPop(t *testing.T) {
	tests := []struct {
		board    Bitboard
		expected int
		after    Bitboard
	}{
		{board: 0x0001, expected: 0, after: 0x0000},
		{board: 0x0003, expected: 1, after: 0x0001},
	}

	for _, test := range tests {
		output := test.board.PopBit()

		if output != test.expected && test.board != test.after {
			t.Errorf("expected sq: %d, output sq: %d, expected bitboard: %x, output bitboard: %x", test.expected, output, test.after, test.board)
		}
	}
}

func TestBitboardCount(t *testing.T) {
	tests := []struct {
		board Bitboard
		expected int
	}{
		{board: 0x0001, expected: 1},
	}

	for _, test := range tests {
		output := test.board.Count()

		if output != test.expected{
			t.Errorf("expected: %d, output: %d", test.expected, output)
		}
	}

}
