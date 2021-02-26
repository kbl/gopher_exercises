package main

import (
	"github.com/kbl/gopher_exercises/book/ch03/surface"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	startServer()
}

func writeToStdout() {
	surface.Write(os.Stdout, surface.Saddle, 600, 320)
}

func startServer() {

	transforms := map[string]surface.ZTransform{
		"saddle": surface.Saddle,
		"peak":   surface.Peak,
		"eggBox": surface.EggBox,
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/svg+xml")
		transform := transforms["peak"]
		width := 600
		height := 320

		for param, values := range r.URL.Query() {
			if param == "transform" {
				t, present := transforms[values[0]]
				if present {
					transform = t
				}
			} else if param == "width" {
				width, _ = strconv.Atoi(values[0])
			} else if param == "height" {
				height, _ = strconv.Atoi(values[0])
			}
		}

		surface.Write(w, transform, width, height)
	})
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
