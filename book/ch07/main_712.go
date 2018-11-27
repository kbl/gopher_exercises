package main

import (
	"book/ch07/ex712"
	"book/ch07/http4"
	"log"
	"net/http"
)

func main() {
	db := new(ex712.Database)
	db.Database.Database = make(map[string]http4.Dollars)
	http.HandleFunc("/list", db.List)
	http.HandleFunc("/price", db.Price)
	http.HandleFunc("/create", db.Create)
	http.HandleFunc("/read", db.Read)
	http.HandleFunc("/update", db.Update)
	http.HandleFunc("/delete", db.Delete)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
