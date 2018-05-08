package mandelbrot

import (
	"image/color"
	"testing"
)

func TestAvgColors(t *testing.T) {
	colors := []color.Color{
		color.RGBA{R: 35, G: 35, B: 220, A: 255},
		color.Black,
		color.Black,
		color.Black,
	}
	expected := color.RGBA{R: 8, G: 8, B: 55, A: 255}

	avg := avgColor(colors)

	if avg != expected {
		t.Errorf("avgColor(%v) = %v, want %v", colors, avg, expected)
	}
}
