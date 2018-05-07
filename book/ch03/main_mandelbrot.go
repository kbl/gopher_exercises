package main

import (
	"book/ch03/mandelbrot"
    "os"
)

func main() {
	mandelbrot.Draw(
        os.Stdout,
        mandelbrot.Mandelbrot128,
    )
	// mandelbrot.Draw(mandelbrot.Newtons)
}
