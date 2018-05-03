package lissajous

import (
    "image"
    "image/color"
    "image/gif"
    "io"
    "math"
    "math/rand"
)

var palette []color.Color = make([]color.Color, 256, 256)

func init() {
    palette[0] = color.Black
    for i := uint8(1); i > 0; i++ {
        palette[i] = color.RGBA{255 - i, i, 255 - i, 255}
    }
}

func Draw(out io.Writer) {
    const (
        cycles  = 5
        res     = 0.001 // angular resolution
        size    = 100   // image canvas covers [-size..+size]
        nframes = 64    // number of animation frames
        delay   = 8     // delay between frames in 10ms units
    )
    freq := rand.Float64() * 3.0 // relative frequency of y oscillator
    anim := gif.GIF{LoopCount: nframes}
    phase := 0.0 // phase difference
    for i := 0; i < nframes; i++ {
        rect := image.Rect(0, 0, 2*size+1, 2*size+1)
        img := image.NewPaletted(rect, palette)
        colorIndex := 0
        for t := 0.0; t < cycles*2*math.Pi; t += res {
            x := math.Sin(t)
            y := math.Sin(t*freq + phase)
            img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
                uint8(colorIndex))
            colorIndex += 1
            colorIndex %= 256
        }
        phase += 0.1
        anim.Delay = append(anim.Delay, delay)
        anim.Image = append(anim.Image, img)
    }
    gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}

