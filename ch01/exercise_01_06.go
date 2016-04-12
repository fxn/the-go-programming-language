/*
  Exercise 1.6: Modify the Lissajous program to produce images in multiple
  colors by adding more values to palette and then displaying them by changing
  the third argument of SetColorIndex in some interesting way.

  I chose some colors from http://raebear.net/comp/emacscolors.html and looked
  up how to generate a random integer. Needed also to explicitly get an uint8
  out of it because the compiler refused to coerce.
*/
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

var palette = []color.Color{
	color.RGBA{0xFF, 0xFF, 0xFF, 0xFF}, // white
	color.RGBA{0x00, 0x00, 0x00, 0xFF}, // black
	color.RGBA{0xFF, 0x00, 0x00, 0xFF}, // red
	color.RGBA{0x00, 0xFF, 0x00, 0xFF}, // green
	color.RGBA{0x00, 0x00, 0xFF, 0xFF}, // blue
	color.RGBA{0xFF, 0xFF, 0x00, 0xFF}, // yellow
	color.RGBA{0x00, 0xFF, 0xFF, 0xFF}, // cyan
	color.RGBA{0xFF, 0x00, 0xFF, 0xFF}, // magenta
}

func main() {
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)

	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference

	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)

		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			i := rand.Intn(len(palette)-1) + 1

			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), uint8(i))
		}

		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}

	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
