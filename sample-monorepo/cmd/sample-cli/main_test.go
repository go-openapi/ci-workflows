package main

import (
	"testing"

	"github.com/go-openapi/testify/v2/assert"
)

func TestGreeting(t *testing.T) {
	assert.Equal(t, "sample-cli:", greeting())
}
