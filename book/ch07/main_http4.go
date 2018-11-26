package main

import (
	"book/ch07/http4"
	"log"
	"net/http"
)

func main() {
	db := http4.Database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.List)
	http.HandleFunc("/price", db.Price)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
