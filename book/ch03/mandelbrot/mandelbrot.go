package mandelbrot

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"math"
	"math/big"
	"math/cmplx"
)

const (
	xmin   = -3
	ymin   = -3
	xmax   = -2
	ymax   = -2
	width  = 1024
	height = 1024
)

type CanvasDetails struct {
	Zoom,
	Xmin,
	Xmax,
	Ymin,
	Ymax float64
}

func (cd CanvasDetails) Xrange() float64 {
	return (cd.Xmax - cd.Xmin) / cd.Zoom
}

func (cd CanvasDetails) Yrange() float64 {
	return (cd.Ymax - cd.Ymin) / cd.Zoom
}

type SetFunction func(complex128) color.Color

var palette [256]color.Color

func init() {
	for i := 0; i < len(palette); i++ {
		gradient := math.Log(float64(i)) / math.Log(float64(len(palette)))
		x := uint8(gradient * float64(len(palette)))
		palette[i] = color.RGBA{R: 255 - x, G: 255 - x, B: x, A: 255}
	}
}

func Draw(out io.Writer, f SetFunction, canvas CanvasDetails) {
	xrange := canvas.Xrange()
	yrange := canvas.Yrange()
	xmin := canvas.Xmin
	ymin := canvas.Ymin

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*yrange + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*xrange + xmin
			c := supersampled(x, y, xrange, yrange, f)
			img.Set(px, py, c)
		}
	}
	png.Encode(out, img)
}

func Mandelbrot128(z complex128) color.Color {
	const iterations = len(palette)

	var v complex128
	for n := 0; n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return palette[n]
		}
	}
	return color.Black
}

func Mandelbrot64(z complex128) color.Color {
	const iterations = len(palette)

	var v complex64
	z64 := complex64(z)

	for n := 0; n < iterations; n++ {
		v = v*v + z64
		if cmplx.Abs(complex128(v)) > 2 {
			return palette[n]
		}
	}
	return color.Black
}

func MandelbrotFloat(z complex128) color.Color {
	const iterations = len(palette)

	var v *ComplexFloat = NewComplexFloat()
	var zFloat *ComplexFloat = &ComplexFloat{
		big.NewFloat(real(z)),
		big.NewFloat(imag(z)),
	}

	for n := 0; n < iterations; n++ {
		v = v.Mul(v, v)
		v = v.Add(v, zFloat)
		if x, _ := v.Abs().Float64(); x > 2 {
			return palette[n]
		}
	}
	return color.Black
}

func Newtons(z complex128) color.Color {
	const iterations = 25
	const tolerance = 0.0001

	roots := []complex128{
		1,
		-1,
		1i,
		-1i,
	}

	rootColors := []color.Color{
		color.RGBA{R: 0, G: 255, B: 255, A: 255},
		color.RGBA{R: 255, G: 255, B: 0, A: 255},
		color.RGBA{R: 255, G: 0, B: 255, A: 255},
		color.RGBA{R: 255, G: 0, B: 0, A: 255},
	}

	shade := func(c color.Color, i int) color.Color {
		var shade uint8 = uint8(float64(i) / float64(iterations) * 255)
		r, g, b, a := c.RGBA()
		r8 := uint8(r)
		g8 := uint8(g)
		b8 := uint8(b)
		if r8 != 0 {
			r8 -= shade
		}
		if g8 != 0 {
			g8 -= shade
		}
		if b8 != 0 {
			b8 -= shade
		}
		return color.RGBA{
			R: r8,
			G: g8,
			B: b8,
			A: uint8(a),
		}
	}

	for n := 0; n < iterations; n++ {
		z -= (cmplx.Pow(z, 4) - 1) / (3 * cmplx.Pow(z, 3))
		for i, root := range roots {
			if cmplx.Abs(root-z) < tolerance {
				return shade(rootColors[i], n)
			}
		}
	}

	return color.White
}

func supersampled(x, y, xrange, yrange float64, f SetFunction) color.Color {
	xDif := 0.5 / width * xrange
	yDif := 0.5 / height * yrange

	var colors []color.Color
	for _, xx := range []float64{-xDif, xDif} {
		for _, yy := range []float64{-yDif, yDif} {
			z := complex(x+xx, y+yy)
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

	return color.RGBA{
		R: uint8(r / uint32(len(colors))),
		G: uint8(g / uint32(len(colors))),
		B: uint8(b / uint32(len(colors))),
		A: uint8(a / uint32(len(colors))),
	}
}
