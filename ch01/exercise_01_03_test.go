/*
  Exercise 1.3: Experiment to measure the difference in running time between our
  potentially inefficient versions and the one that uses strings.Join. (Section
  1.6 illustrates part of the time package, and Section 11.4 shows how to write
  benchamrk tests for systematic performance evaluation.

  Solution: I solved this by cargo-culting section 11.4 as suggested. Checked
  slice literals in the official docs.

  echo3 is about twice times faster than echo1.

  By passing -benchmem I could also see that echo3 allocates half as much:
  2 allocs/op, versus 4 allocs/op of echo1.
*/
package exercise_01_03_test

import (
	"strings"
	"testing"
)

var args = []string{"/usr/local/bin/go", "foo", "bar", "baz", "zoo", "woo"}

func BenchmarkEcho1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var _, sep string
		for j := 1; j < len(args); j++ {
			_ += sep + args[j]
			sep = " "
		}
	}
}

func BenchmarkEcho3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strings.Join(args[1:], " ")
	}
}
