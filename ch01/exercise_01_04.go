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
	lines2fnames := make(map[string]map[string]int)
	fnames := os.Args[1:]

	if len(fnames) == 0 {
		countLines("stdin", os.Stdin, lines2fnames)
	} else {
		countLinesForFnames(fnames, lines2fnames)
	}
	printDups(lines2fnames)
}

func countLines(fname string, f *os.File, lines2fnames map[string]map[string]int) {
	input := bufio.NewScanner(f)

	for input.Scan() {
		line := input.Text()
		if len(lines2fnames[line]) == 0 {
			lines2fnames[line] = make(map[string]int)
		}
		lines2fnames[line][fname]++
	}
	// Ignoring potential errors from input.Err(), as in the original dup2.
}

func countLinesForFnames(fnames []string, lines2fnames map[string]map[string]int) {
	for _, fname := range fnames {
		f, err := os.Open(fname)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
			continue
		}
		countLines(fname, f, lines2fnames)
		f.Close()
	}
}

func printDups(lines2fnames map[string]map[string]int) {
	for line, files := range lines2fnames {
		total := 0
		for _, count := range files {
			total += count
		}
		if total > 1 {
			fmt.Printf("%d\t%s\n", total, line)
			for fname, _ := range files {
				fmt.Println(fname)
			}			
		}
	}
}