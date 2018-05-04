package surface

import (
    "fmt"
    "math"
    "io"
)

const (
    cells      = 100
    xyrange    = 30.0
    angle      = math.Pi / 6
    colorRange = 4 * 256
)

type drawingDetails struct {
    width,
    height,
    xyscale,
    zscale  float64
}

var sin30, cos30 = math.Sincos(angle)

type ZTransform func(float64, float64) float64

func Write(out io.Writer, zTransform ZTransform, width, height int) {
    params := drawingDetails{
        width:   float64(width),
        height:  float64(height),
        xyscale: float64(width) / 2 / xyrange,
        zscale:  float64(height) * 0.4,
    }

    fmt.Fprintf(out, "<svg xmlns='http://www.w3.org/2000/svg' " +
               "style='stroke: grey; fill: white; stroke-width: 0.7' " +
               "width='%d' height='%d'>", params.width, params.height)

    zMin, zMax := findZRange(zTransform)

    for i := 0; i < cells; i++ {
        for j := 0; j < cells; j++ {
            ax, ay, az := corner(i + 1, j,     params, zTransform)
            bx, by, bz := corner(i,     j,     params, zTransform)
            cx, cy, cz := corner(i,     j + 1, params, zTransform)
            dx, dy, dz := corner(i + 1, j + 1, params, zTransform)

            color := findColor(zMin, zMax, (az + bz + cz + dz) / 4)

            if math.IsNaN(ax + ay + bx + by + cx + cy + dx + dy) {
                continue
            }

            fmt.Fprintf(out, "<polygon points='%g,%g %g,%g %g,%g %g,%g' fill='%s'/>\n", ax, ay, bx, by, cx, cy, dx, dy, color)
        }
    }
    fmt.Fprintf(out, "</svg>\n")
}

func corner(i, j int, params drawingDetails, zTransform ZTransform) (float64, float64, float64) {
    x := xyrange * (float64(i) / cells - 0.5)
    y := xyrange * (float64(j) / cells - 0.5)

    z := zTransform(x, y)

    sx := params.width / 2 + (x - y) * cos30 * params.xyscale
    sy := params.height / 2 + (x + y) * sin30 * params.xyscale - z * params.zscale

    return sx, sy, z
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

func findColor(zMin, zMax, z float64) string {
    z = (z - zMin) / (zMax - zMin)
    color := int(math.Ceil(z * colorRange))
    if color <= 255 {
        return fmt.Sprintf("rgb(0, 0, %d)", color)
    }
    if color <= 255 * 2 {
        color = color - 255
        return fmt.Sprintf("rgb(0, %d, 255)", color)
    }
    if color <= 255 * 3 {
        color = color - 255 * 2
        return fmt.Sprintf("rgb(0, 255, %d)", 255 - color)
    }
    if color <= 255 * 4 {
        color = color - 255 * 3
        return fmt.Sprintf("rgb(%d, 255, 0)", color)
    }
    color = color - 255 * 4
    return fmt.Sprintf("rgb(255, %d, 0)", 255 - color)
}

func findZRange(zTransform ZTransform) (float64, float64) {
    zMin := math.MaxFloat64
    zMax := -math.MaxFloat64
    for x := -xyrange / 2; x <= xyrange / 2; x += 0.5 {
        for y := -xyrange / 2; y <= xyrange / 2; y += 0.5 {
            value := zTransform(x, y)
            if math.IsNaN(value) {
                continue
            }
            if value > zMax {
                zMax = value
            }
            if value < zMin {
                zMin = value
            }
        }
    }
    return zMin, zMax
}
