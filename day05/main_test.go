package main

import "testing"

func TestDecode(t *testing.T) {
	tests := []struct {
		arg      string
		row, col int
	}{
		{
			arg: "FBFBBFFRLR",
			row: 44,
			col: 5,
		},
		{
			arg: "FFFBBBFRRR",
			row: 14,
			col: 7,
		},
		{
			arg: "BBFFBBFRLL",
			row: 102,
			col: 4,
		},
	}

	for _, tt := range tests {
		r, c := decode(tt.arg)
		if r != tt.row {
			t.Errorf("row error: want %d, got %d", tt.row, r)
		}
		if c != tt.col {
			t.Errorf("row error: want %d, got %d", tt.col, c)
		}
	}

}
