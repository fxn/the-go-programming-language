/*
  Exercise 1.11: Try fetchall with longer argument lists, such as samples from
  the top million web sites available at alexa.com. How does the program behave
  if a web site just doesn't respond? (Section 8.9 describes mechanisms for
  coping in such cases.)

  Solution: The top million web sites is available from Alexa in this ZIP:

      http://s3.amazonaws.com/alexa-static/top-1m.csv.zip

  I grepped /usr/include/limits.h for ARG_MAX and saw _POSIX_ARG_MAX, defined to
  be 4096, so I passed those many to fetchall:

      head -n 4096 top-1m.csv | \
      cut -d, -f2 | \
      perl -pe 's!^!http://!' | \
      xargs ./fetchall

  The entire thing took about 30s. About 35 of the web sites didn't respond. In
  that case the goroutine gets an error from http.Get, that is passed to the
  channel as a string, as can be seen in the listing below. The error message
  looks like this:

      Get http://ticketmaster.com: dial tcp: i/o timeout

  I copy here fetchall.go for reference.
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
	start := time.Now()
	ch := make(chan string)

	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}

	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
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
