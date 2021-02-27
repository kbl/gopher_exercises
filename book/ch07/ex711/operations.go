package ex711

import (
	"fmt"
	"github.com/kbl/gopher_exercises/book/ch07/http4"
	"net/http"
	"strconv"
)

type Database struct {
	http4.Database
}

func (db Database) respondIfMissing(w http.ResponseWriter, item string) bool {
	_, ok := db.Database[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
	return !ok
}

func (db Database) respondIfExist(w http.ResponseWriter, item string) bool {
	_, ok := db.Database[item]
	if ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "item already exists: %q\n", item)
	}
	return ok
}

func (db Database) Create(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("name")
	if db.respondIfExist(w, item) {
		return
	}
	priceReq := req.URL.Query().Get("price")
	price, err := strconv.ParseFloat(priceReq, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "price %q is invalid!\n", priceReq)
		return
	}
	db.Database[item] = http4.Dollars(price)
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "item created: %q\n", item)
}

func (db Database) Read(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("name")
	if db.respondIfMissing(w, item) {
		return
	}
	price, _ := db.Database[item]
	fmt.Fprintf(w, "%q: %q\n", item, price)
}

func (db Database) Update(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("name")
	if db.respondIfMissing(w, item) {
		return
	}
	priceReq := req.URL.Query().Get("price")
	price, err := strconv.ParseFloat(priceReq, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "price %q is invalid!\n", priceReq)
		return
	}
	db.Database[item] = http4.Dollars(price)
}

func (db Database) Delete(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("name")
	if db.respondIfMissing(w, item) {
		return
	}
	delete(db.Database, item)
}
