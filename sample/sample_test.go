package sample

import (
	"testing"

	"github.com/go-openapi/testify/v2/require"
)

func TestSample(t *testing.T) {
	/*
		if runtime.GOOS == "windows" {
			require.Equal(t, 0, Sample()) // make the test fail on windows

			return
		}

	*/
	require.Equal(t, 1, Sample())
}

func FuzzSample(f *testing.F) {
	f.Add([]byte(nil))
	f.Add([]byte{})
	f.Add([]byte{'x'})

	f.Fuzz(func(t *testing.T, input []byte) {
		require.NotPanics(t, func() {
			_ = fuzzable(input)
		})
	})
}
