package main

import "testing"

func TestFileValid(t *testing.T) {
	tests := []struct {
		file     File
		expected bool
	}{
		{file: 1, expected: true},
		{file: 7, expected: true},
		{file: -1, expected: false},
		{file: 10, expected: false},
	}

	for _, test := range tests {
		output := test.file.Valid()

		if output != test.expected {
			t.Errorf("expected: %t, output: %t", test.expected, output)
		}

	}

}
