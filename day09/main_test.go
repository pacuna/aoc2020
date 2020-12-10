package main

import "testing"

func TestHasSum(t *testing.T) {
	tests := []struct {
		in     []int
		target int
		out    bool
	}{
		{
			in:     []int{35, 20, 15, 25, 47},
			target: 40,
			out:    true,
		},
		{
			in:     []int{20, 15, 25, 47, 40},
			target: 62,
			out:    true,
		},
		{
			in:     []int{95, 102, 117, 150, 182},
			target: 127,
			out:    false,
		},
	}

	for _, tt := range tests {
		if res := hasSum(tt.in, tt.target); res != tt.out {
			t.Errorf("error: wanted %v, got %v instead", tt.out, res)
		}
	}
}
