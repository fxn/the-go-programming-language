/*
  Exercise 1.8: Modify fetch to add the prefix http:// to each argument URL if
  it is missing. You might want to use strings.HasPrefix.

  I have used the variant in exercise 1.7.
*/

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

const url_prefix = "http://"

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, url_prefix) {
			url = url_prefix + url
		}

		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		_, err = io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: copying: %v\n", err)
			os.Exit(1)
		}
	}
}
