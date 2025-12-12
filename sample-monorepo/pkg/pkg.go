// Package pkg exercises CI pipelines.
package pkg

func Pkg() string {
	return ""
}

func fuzzable(input []byte) string {
	if len(input) > 0 {
		return string(input)
	}

	return "0"
}
