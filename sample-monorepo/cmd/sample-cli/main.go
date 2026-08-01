// Command sample-cli exercises CI pipelines for a mono-repo module distributed as a binary.
//
// Unlike the library modules of this repo, this module carries no "replace" directive for its
// intra-repo dependencies: "go install" refuses any module whose go.mod contains one.
// The release process must therefore resolve its cross-module dependencies without them.
package main

import (
	"fmt"
	"os"

	"github.com/go-openapi/ci-workflows/sample-monorepo/pkg"
)

func main() {
	fmt.Fprintln(os.Stdout, greeting())
}

func greeting() string {
	return "sample-cli:" + pkg.Pkg()
}
