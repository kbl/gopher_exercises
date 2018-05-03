package surface

import (
    "fmt"
    "math"
    "io"
    "os"
)

const (
    width, height = 600, 320
    cells = 100
    xyrange = 30.0
    xyscale = width / 2 / xyrange
    zscale = height * 0.4
    angle = math.Pi / 6
)

var sin30, cos30 = math.Sincos(angle)

type ZTransform func(float64, float64) float64

func Write(out io.Writer, zTransform ZTransform) {
    fmt.Fprintf(out, "<svg xmlns='http://www.w3.org/2000/svg' " +
               "style='stroke: grey; fill: white; stroke-width: 0.7' " +
               "width='%d' height='%d'>", width, height)
    for i := 0; i < cells; i++ {
        for j := 0; j < cells; j++ {
            ax, ay := corner(i + 1, j, zTransform)
            bx, by := corner(i, j, zTransform)
            cx, cy := corner(i, j + 1, zTransform)
            dx, dy := corner(i + 1, j + 1, zTransform)

            if math.IsNaN(ax + ay + bx + by + cx + cy + dx + dy) {
                continue
            }

            fmt.Fprintf(out, "<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n", ax, ay, bx, by, cx, cy, dx, dy)
        }
    }
    fmt.Fprintf(out, "</svg>\n")
}

func corner(i, j int, zTransform ZTransform) (float64, float64) {
    x := xyrange * (float64(i) / cells - 0.5)
    y := xyrange * (float64(j) / cells - 0.5)

    z := zTransform(x, y)

    sx := width / 2 + (x - y) * cos30 * xyscale
    sy := height / 2 + (x + y) * sin30 * xyscale - z * zscale

    return sx, sy
}

func Peak(x, y float64) float64 {
    r := math.Hypot(x, y)
    return math.Sin(r) / r
}

func Saddle(x, y float64) float64 {
    xaxis := -0.006 * x * x
    yaxis := 0.001 * y * y
    return xaxis + yaxis
}

func EggBox(x, y float64) float64 {
    return 0.1 * (math.Cos(x) + math.Cos(y))
}

