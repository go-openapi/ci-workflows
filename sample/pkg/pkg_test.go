package pkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPkg(t *testing.T) {
	assert.Empty(t, Pkg())
}
