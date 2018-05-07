package main

import (
	"book/ch03/mandelbrot"
)

func main() {
	mandelbrot.Draw(mandelbrot.MandelbrotFloat)
	// mandelbrot.Draw(mandelbrot.Newtons)
}
