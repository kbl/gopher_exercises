package mandelbrot

import (
    "image"
    "image/color"
    "image/png"
    "math"
    "math/cmplx"
    "os"
    "math/big"
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

func Mandelbrot128(z complex128) color.Color {
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

func Mandelbrot64(z complex128) color.Color {
    const iterations = len(palette)

    var v complex64
    z64 := complex64(z)

    for n := 0; n < iterations; n++ {
        v = v * v + z64
        if cmplx.Abs(complex128(v)) > 2 {
            return palette[n]
        }
    }
    return color.Black
}

type ComplexFloat struct {
    r, i *big.Float
}

func NewComplexFloat() *ComplexFloat {
    cf := new(ComplexFloat)
    cf.r = new(big.Float)
    cf.i = new(big.Float)
    return cf
}

func (z *ComplexFloat) Add(x, y *ComplexFloat) *ComplexFloat {
    var r, i *big.Float = z.r, z.i

    r.Add(x.r, y.r)
    i.Add(x.i, y.i)

    return z
}

func (z *ComplexFloat) Mul(x, y *ComplexFloat) *ComplexFloat {
    var newR1 *big.Float = &big.Float{}
    var newR2 *big.Float = &big.Float{}
    var newI1 *big.Float = &big.Float{}
    var newI2 *big.Float = &big.Float{}

    var r, i *big.Float = z.r, z.i

    newR1 = newR1.Mul(x.r, y.r)
    newR2.Mul(x.i, y.i)
    r.Sub(newR1, newR2)

    newI1.Mul(x.r, y.i)
    newI2.Mul(x.i, y.r)
    i.Add(newI1, newI2)

    return z
}

func (c *ComplexFloat) Abs() *big.Float {
    var f1 *big.Float = &big.Float{}
    var f2 *big.Float = &big.Float{}

    f1.Mul(c.r, c.r)
    f2.Mul(c.i, c.i)
    f1.Add(f1, f2)

    return f1.Sqrt(f1)
}

func MandelbrotFloat(z complex128) color.Color {
    const iterations = len(palette)

    var v *ComplexFloat = NewComplexFloat()
    var zFloat *ComplexFloat = &ComplexFloat {big.NewFloat(real(z)), big.NewFloat(imag(z))}

    for n := 0; n < iterations; n++ {
        v = v.Mul(v, v)
        v = v.Add(v, zFloat)
        if x, _ := v.Abs().Float64(); x > 2 {
            return palette[n]
        }
    }
    return color.Black
}


// type ComplexRat struct {
//     r, i big.Rat
// }
//
// func (c1 ComplexRat) Add(c2 ComplexRat) ComplexRat {
//     return ComplexRat {
//         c1.r + c2.r,
//         c1.i + c2.i,
//     }
// }
// 
// func (c1 ComplexRat) Mul(c2 ComplexRat) ComplexRat {
//     return ComplexRat {
//         c1.r * c2.r - c1.i * c2.i,
//         c1.r * c2.i + c1.i * c2.r,
//     }
// }
// 
// func (c ComplexRat) Abs() big.Rat {
//     return sqrt(c.r * c.r + c.i * c.i)
// }


func Newtons(z complex128) color.Color {
    const iterations = 25
    const tolerance = 0.0001

    roots := []complex128 {
        1,
        -1,
        1i,
        -1i,
    }

    rootColors := []color.Color {
        color.RGBA {R: 0,   G: 255, B: 255, A: 255},
        color.RGBA {R: 255, G: 255, B: 0,   A: 255},
        color.RGBA {R: 255, G: 0,   B: 255, A: 255},
        color.RGBA {R: 255, G: 0,   B: 0,   A: 255},
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
        return color.RGBA {
            R: r8,
            G: g8,
            B: b8,
            A: uint8(a),
        }
    }

    for n := 0; n < iterations; n++ {
        z -= (cmplx.Pow(z, 4) - 1) / (3 * cmplx.Pow(z, 3))
        for i, root := range roots {
            if cmplx.Abs(root - z) < tolerance {
                return shade(rootColors[i], n)
            }
        }
    }

    return color.White
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
