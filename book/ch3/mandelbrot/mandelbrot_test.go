package mandelbrot

import (
    "testing"
    "image/color"
    "fmt"
)

func TestAvgColors(t *testing.T) {
    colors := []color.Color {
        color.RGBA {R: uint8(35), G: uint8(35), B: uint8(220), A: uint8(255)},
        color.Black,
        color.Black,
        color.Black,
    }
    expected := color.RGBA{R: uint8(8), G: uint8(8), B: uint8(55), A: uint8(255)}

    avg := avgColor(colors)

    if avg != expected {
        t.Errorf("avgColor(%v) = %v, want %v", colors, avg, expected)
    }
}
