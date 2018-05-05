package mandelbrot

import (
    "image"
    "image/color"
    "image/png"
    "math"
    "math/cmplx"
    "os"
    "fmt"
)

const (
    xmin   = -2
    ymin   = -2
    xmax   = 2
    ymax   = 2
    width  = 1024
    height = 1024
    xrange = xmax - xmin
    yrange = ymax - ymin
)

type SetFunction func(complex128) color.Color

var palette [256]color.Color

func init() {
    for i := 0; i < len(palette); i++ {
        gradient := math.Log(float64(i)) / math.Log(float64(len(palette)))
        x := uint8(gradient * float64(len(palette)))
        palette[i] = color.RGBA {R: 255 - x, G: 255 - x, B: x, A: 255}
    }
}

func Draw(f SetFunction) {
    img := image.NewRGBA(image.Rect(0, 0, width, height))
    for py := 0; py < height; py++ {
        y := float64(py) / height * yrange + ymin
        for px := 0; px < width; px++ {
            x := float64(px) / width * xrange + xmin
            c := supersampled(x, y, f)
            img.Set(px, py, c)
        }
    }
    png.Encode(os.Stdout, img)
}

func Mandelbrot(z complex128) color.Color {
    const iterations = len(palette)

    roots := []complex128 {
        1,
        -1,
        1i,
        -1i
    }

    rootColors := []color.Color {
        // Cyan	(1,0,0,0)	#00FFFF
        color.RGBA {R: uint8(255), G: uint8(255), B: 0, uint8(255)},
        // Magenta	(0,1,0,0)	#FF00FF
        color.Black
    }

    fmt.Println(roots, rootColors)
    return color.Black
}

func Newtons(z complex128) color.Color {
    const iterations = len(palette)

    var v complex128
    for n := 0; n < iterations; n++ {
        v = v * v * v * v + z
        if cmplx.Abs(v) > 2 {
            return palette[n]
        }
    }
    return color.Black
}

func newtons(z complex128) color.Color {
    const iterations = len(palette)

    var v complex128
    for n := 0; n < iterations; n++ {
        v = v * v + z
        if cmplx.Abs(v) > 2 {
            return palette[n]
        }
    }
    return color.Black
}

func supersampled(x, y float64, f SetFunction) color.Color {
    xDif := 0.5 / width * xrange
    yDif := 0.5 / height * yrange

    var colors []color.Color
    for _, xx := range []float64 {-xDif, xDif} {
        for _, yy := range []float64 {-yDif, yDif} {
            z := complex(x + xx, y + yy)
            c := f(z)
            colors = append(colors, c)
        }
    }

    return avgColor(colors)
}

func avgColor(colors []color.Color) color.Color {
    var r, g, b, a uint32

    for _, c := range colors {
        rTemp, gTemp, bTemp, aTemp := c.RGBA()
        r += uint32(uint8(rTemp))
        g += uint32(uint8(gTemp))
        b += uint32(uint8(bTemp))
        a += uint32(uint8(aTemp))
    }

    return color.RGBA {
        R: uint8(r / uint32(len(colors))),
        G: uint8(g / uint32(len(colors))),
        B: uint8(b / uint32(len(colors))),
        A: uint8(a / uint32(len(colors))),
    }
}
