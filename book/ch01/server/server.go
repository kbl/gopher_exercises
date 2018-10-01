package server

import (
	"book/ch01/lissajous"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

var mu sync.Mutex
var count int

func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Println("invoked /")
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}

func Start() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	http.HandleFunc("/lisa", func(w http.ResponseWriter, r *http.Request) {
		var cycles int = 5
		for param, values := range r.URL.Query() {
			if param == "count" {
				cycles, _ = strconv.Atoi(values[0])
			}
		}
		lissajous.Draw(cycles, w)
	})
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
