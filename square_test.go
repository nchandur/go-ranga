package main

import "testing"

func TestSquareValid(t *testing.T) {
	tests := []struct {
		square      Square
		expected    bool
	}{
		{square: A1, expected: true},
		{square: NoSquare, expected: true},
		{square: 1000, expected: false},
		{square: -10, expected: false},
	}

	for _, test := range tests {

		output := test.square.Valid()

		if output != test.expected {
			t.Errorf("expected %t, output: %t", test.expected, output)
		}

	}

}
