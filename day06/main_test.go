package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSolveI(t *testing.T) {
	res, err := solveI("example.txt")
	require.NoError(t, err)
	require.Equal(t, 11, res)
}

func TestSolveII(t *testing.T) {
	res, err := solveII("example.txt")
	require.NoError(t, err)
	require.Equal(t, 6, res)
}
