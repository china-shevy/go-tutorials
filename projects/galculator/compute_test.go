package main

import "testing"
import "github.com/stretchr/testify/require"

func Test(t *testing.T) {
	require.Equal(t, "-4", compute("1 + 1 - 3 * 9 - 1 / 3"))
	require.Equal(t, "-8", compute("1 + 1 - 10"))
}
