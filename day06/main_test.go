package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSolve(t *testing.T) {
	res, err := solve("example.txt")
	require.NoError(t, err)
	require.Equal(t, 11, res)
}
