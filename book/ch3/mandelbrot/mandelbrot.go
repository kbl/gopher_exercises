package mandelbrot

import (
    "image"
    "image/color"
    "image/png"
    "math"
    "math/cmplx"
    "os"
)

const (
    xmin   = -2
    ymin   = -2
    xmax   = 2
    ymax   = 2
    width  = 1024
    height = 1024
)

var palette [256]color.Color

func init() {
    for i := 0; i < len(palette); i++ {
        gradient := math.Log(float64(i)) / math.Log(float64(len(palette)))
        x := uint8(gradient * float64(len(palette)))
        palette[i] = color.RGBA{R: 255 - x, G: 255 - x, B: x, A: 255}
    }
}

func Draw() {
    img := image.NewRGBA(image.Rect(0, 0, width, height))
    for py := 0; py < height; py++ {
        y := float64(py) / height * (ymax - ymin) + ymin
        for px := 0; px < width; px++ {
            x := float64(px) / width * (xmax - xmin) + xmin
            z := complex(x, y)
            img.Set(px, py, mandelbrot(z))
        }
    }
    png.Encode(os.Stdout, img)
}

func mandelbrot(z complex128) color.Color {
    const iterations = 200

    var v complex128
    for n := 0; n < iterations; n++ {
        v = v * v + z
        if cmplx.Abs(v) > 2 {
            return palette[n]
        }
    }
    return color.Black
}
