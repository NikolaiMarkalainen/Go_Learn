package main

import (
	"os"
	"strings"
	"testing"
)

// 1.3 benchmarks
func BenchmarkForLoop(b *testing.B) {
	var input, sep string
	for i := 0; i < b.N; i++ {
		for _, char := range os.Args[1:len(os.Args)] {
			// 1.2	os.Args[0 +1 onwards]
			input += sep + char
			sep = " "
		}
	}
}
func BenchmarkStringJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strings.Join(os.Args[1:], " ")

	}
}
