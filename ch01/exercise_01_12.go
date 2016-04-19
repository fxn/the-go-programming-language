/*
  Exercise 1.12: Modify the Lissajous server to read parameter values from the
  URL. For example, you might arrange it so that a URL link

    http://localhost:8000/?cycles=20

  sets the number of cycles to 20 instead of the default 5. Use the strconv.Atoi
  function to convert the string parameter into an integer. You can see its
  documentation with go doc strconv.Atoi.

  Solution: I extracted the values that parametrized the GIF into variables
  initialized to the same literals the constants are in the original listing.
  The code needs to convert some values to be able to do arithmetic with them.
  For example, I wrote `float64(2*cycles)*math.Pi` because with the vanilla
  product the compiler complained: "constant 3.14159 truncated to integer".

  Do not know yet which is the idiomatic way to deal with these type mismatches,
  and which are the actual type-related rules in these expressions. Guess they
  will be covered later in the book.
*/
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
)

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0
	blackIndex = 1
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	cycles := 5   // number of complete x oscillator revolutions
	res := 0.001  // angular resolution
	size := 100   // image canvas covers [-size..+size]
	nframes := 64 // number of animation frames
	delay := 8    // delay between frames in 10ms units

	if err := r.ParseForm; err != nil {
		log.Print(err)
	}

	for param, values := range r.Form {
		value := values[0]

		// We are going to convert existing parameters only if the conversion
		// succeeds. Otherwise just ignore them.
		switch param {
		case "cycles":
			if i, err := strconv.Atoi(value); err == nil {
				cycles = i
			}
		case "res":
			if f, err := strconv.ParseFloat(value, 64); err == nil {
				res = f
			}
		case "size":
			if i, err := strconv.Atoi(value); err == nil {
				size = i
			}
		case "nframes":
			if i, err := strconv.Atoi(value); err == nil {
				nframes = i
			}
		case "delay":
			if i, err := strconv.Atoi(value); err == nil {
				delay = i
			}
		}
	}

	lissajous(w, cycles, res, size, nframes, delay)
}

func lissajous(out io.Writer, cycles int, res float64, size int, nframes int, delay int) {
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference

	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)

		for t := 0.0; t < float64(2*cycles)*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)

			img.SetColorIndex(size+int(x*float64(size)+0.5), size+int(y*float64(size)+0.5), blackIndex)
		}

		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}

	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
