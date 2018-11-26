package http4

import (
	"fmt"
	"net/http"
)

type Dollars float32

func (d Dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

type Database map[string]Dollars

func (db Database) List(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%q: %q\n", item, price)
	}
}

func (db Database) Price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%q\n", price)
}
