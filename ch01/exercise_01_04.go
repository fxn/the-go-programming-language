/*
  Exercise 1.4: Modify dup2 to print the names of all files in which each
  duplicated line occurs.

  Solution: I have edited a little bit the original listing, some variable
  renamed, blank lines added, code extracted into functions, etc.

  Since the counter in dup2 is not per file, but global, the program needs to
  keep track of the origin of each line. Reason is, if a line appears first in
  file F, and only repeated later in file G, the program still needs to report
  F and G for that line.

  Had to investigate how to declare and initialize nested maps. The zero value
  for a map is printed by %v as an empty map, but smw_ explained in #go-nuts
  that they need explicit initialization with make() as seen in countLines().

  Also, had to lookup online how to check for key existence.
*/
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	lines2fnames := make(map[string]map[string]bool)
	fnames := os.Args[1:]

	if len(fnames) == 0 {
		countLines("stdin", os.Stdin, counts, lines2fnames)
	} else {
		countLinesForFnames(fnames, counts, lines2fnames)
	}
	printDups(counts, lines2fnames)
}

func countLines(fname string, f *os.File, counts map[string]int, lines2fnames map[string]map[string]bool) {
	input := bufio.NewScanner(f)

	for input.Scan() {
		line := input.Text()

		counts[line]++

		_, ok := lines2fnames[line]
		if !ok {
			lines2fnames[line] = make(map[string]bool)
		}

		lines2fnames[line][fname] = true
	}
	// Ignoring potential errors from input.Err(), as in the original dup2.
}

func countLinesForFnames(fnames []string, counts map[string]int, lines2fnames map[string]map[string]bool) {
	for _, fname := range fnames {
		f, err := os.Open(fname)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
			continue
		}
		countLines(fname, f, counts, lines2fnames)
		f.Close()
	}
}

func printDups(counts map[string]int, lines2fnames map[string]map[string]bool) {
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
			for fname := range lines2fnames[line] {
				fmt.Println(fname)
			}
		}
	}
}
