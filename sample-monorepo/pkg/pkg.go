// Package pkg exercises CI pipelines.
package pkg

import (
	rootpkg "github.com/go-openapi/ci-workflows/sample/pkg"
)

// Pkg depends on the root module of this repo, so that the release process
// has an intra-repo dependency to rewrite.
func Pkg() string {
	return rootpkg.Pkg()
}

func fuzzable(input []byte) string {
	if len(input) > 0 {
		return string(input)
	}

	return "0"
}
