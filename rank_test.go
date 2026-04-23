package main

import "testing"

func TestRankValid(t *testing.T) {
	tests := []struct {
		rank     Rank
		expected bool
	}{
		{rank: 1, expected: true},
		{rank: 7, expected: true},
		{rank: -1, expected: false},
		{rank: 10, expected: false},
	}

	for _, test := range tests {
		output := test.rank.Valid()

		if output != test.expected {
			t.Errorf("expected: %t, output: %t", test.expected, output)
		}

	}

}
