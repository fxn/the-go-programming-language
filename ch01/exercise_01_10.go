/*
  Exercise 1.10: Find a web site that produces a large amount of data.
  Investigate caching by running fetchall twice in succession to see whether the
  reported time changes much. Do you get the same content each time? Modify
  fetchall to print its output to a file so it can be examined.

  Solution: For this exercise I chose Rails Contributors

      http://contributors.rubyonrails.org

  which I wrote myself. This web site has a somewhat big table in its front
  page, and it is cached using page caching.

  I expired manually the cache in the server and fetchall reported +3s in the
  first run. A second run took about 0.70s. The served content does not differ.

  I chose to pass the file in which to store the output as an argument. Googled
  how to write to a file.
*/
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintf(os.Stderr, "usage: ./exercise_01_10 filename\n")
		os.Exit(1)
	}

	filename := os.Args[1]
	urls := os.Args[2:]

	out, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer out.Close()

	start := time.Now()
	ch := make(chan string)

	for _, url := range urls {
		go fetch(url, ch)
	}

	for range urls {
		fmt.Fprintln(out, <-ch)
	}
	fmt.Fprintf(out, "%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}
