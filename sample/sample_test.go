package sample

import (
	"testing"

	"github.com/stretchr/testify/require"
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
