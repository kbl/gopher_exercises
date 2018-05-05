package main

import (
	"book/ch3/mandelbrot"
)

func main() {
	mandelbrot.Draw(mandelbrot.MandelbrotFloat)
	// mandelbrot.Draw(mandelbrot.Newtons)
}
