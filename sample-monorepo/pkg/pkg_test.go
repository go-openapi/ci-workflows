package pkg

import (
	"testing"

	"github.com/go-openapi/testify/v2/assert"
	"github.com/go-openapi/testify/v2/require"
)

func TestPkg(t *testing.T) {
	assert.Empty(t, Pkg())
}

func FuzzMonorepo(f *testing.F) {
	f.Add([]byte(nil))
	f.Add([]byte{})
	f.Add([]byte{'x'})

	f.Fuzz(func(t *testing.T, input []byte) {
		require.NotPanics(t, func() {
			_ = fuzzable(input)
		})
	})
}
