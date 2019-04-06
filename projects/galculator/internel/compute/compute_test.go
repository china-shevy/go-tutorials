package compute

import "testing"
import "github.com/stretchr/testify/require"

func TestCompute(t *testing.T) {
	require.Equal(t, "-3", Compute("1 + 1 - 3 * 9 - 1 / 3"))
	require.Equal(t, "-8", Compute("1 + 1 - 10"))
	require.Equal(t, "2", Compute("1+1"))
}
