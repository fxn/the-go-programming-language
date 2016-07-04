/*
  Exercise 1.2: Modify the echo program to print the index and value of each of
  its arguments, one per line.

  Solution: This program is based on echo2. Note that fmt.Println accepts a
  variable number of arguments and adds a space between them.

  See https://golang.org/pkg/fmt/#Println.
*/
package main

import (
	"fmt"
	"os"
)

func main() {
	for i, arg := range os.Args[1:] {
		fmt.Println(i, arg)
	}
}
