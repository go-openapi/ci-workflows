// Package sample is used to exercise CI pipelines.
package sample

func Sample() int {
	return 1
}

func fuzzable(input []byte) string {
	if len(input) > 0 {
		return string(input)
	}

	return "0"
}
