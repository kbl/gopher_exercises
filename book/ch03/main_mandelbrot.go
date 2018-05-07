package main

import (
	"book/ch03/mandelbrot"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/png")
		zoom := 1.0
		xmin := -2.0
		ymin := -2.0
		xmax := 2.0
		ymax := 2.0

		for param, values := range r.URL.Query() {
			if param == "zoom" {
				zoom, _ = strconv.ParseFloat(values[0], 64)
			} else if param == "xmin" {
				xmin, _ = strconv.ParseFloat(values[0], 64)
			} else if param == "xmax" {
				xmax, _ = strconv.ParseFloat(values[0], 64)
			} else if param == "ymin" {
				ymin, _ = strconv.ParseFloat(values[0], 64)
			} else if param == "ymax" {
				ymax, _ = strconv.ParseFloat(values[0], 64)
			}
		}

		mandelbrot.Draw(
			w,
			mandelbrot.Mandelbrot128,
			mandelbrot.CanvasDetails{
				zoom,
				xmin,
				xmax,
				ymin,
				ymax,
			},
		)
	})
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
