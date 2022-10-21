package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)


func TestPart1(t *testing.T) {
	require.Equal(t, 0, part1(1))
	require.Equal(t, 12, part1(3))
	require.Equal(t, 23, part1(2))
	require.Equal(t, 1024, part1(31))
}
