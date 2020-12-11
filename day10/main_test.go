package main

import (
	"strings"
	"testing"
)

func TestSolveI(t *testing.T) {
	input := strings.NewReader(`16
10
15
5
1
11
7
19
6
12
4`)

	d1, d3 := solveI(input)
	if d1 != 7 {
		t.Errorf("should be 7, got %d", d1)
	}
	if d3 != 5 {
		t.Errorf("should be 5, got %d", d3)
	}

}

func TestSolveI_II(t *testing.T) {
	input := strings.NewReader(`28
33
18
42
31
14
46
20
48
47
24
23
49
45
19
38
39
11
1
32
25
35
8
17
7
9
4
2
34
10
3`)

	d1, d3 := solveI(input)
	if d1 != 22 {
		t.Errorf("should be 22, got %d", d1)
	}
	if d3 != 10 {
		t.Errorf("should be 10, got %d", d3)
	}
}
