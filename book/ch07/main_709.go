package main

import (
	"book/ch07/ex709"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		for param, values := range r.URL.Query() {
			if param == "s" {
				ex709.TT.Sort(values[0])
			}
		}
		ex709.HTMLTemplate.Execute(w, ex709.TT)
	})
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
