/*
  Exercise 1.9: Modify fetch to also print the HTTP status code, found in
  resp.Status.

  I have used the variant in exercise 1.7.

  Turns out resp.Status is a string like "200 OK", whereas resp.StatusCode is an
  actual integer with the code (200).
*/
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("HTTP status: %s\n", resp.Status)

		_, err = io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: copying: %v\n", err)
			os.Exit(1)
		}
	}
}
