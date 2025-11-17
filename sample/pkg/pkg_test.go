package pkg

import (
	"testing"

	"github.com/go-openapi/testify/v2/assert"
)

func TestPkg(t *testing.T) {
	assert.Empty(t, Pkg())
}
