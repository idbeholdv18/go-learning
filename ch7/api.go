package main

import (
	"go-learning/ch7/api"
	"log"
	"net/http"
)

func main() {
	db := api.Database{"shoes": 30, "socks": 5}

	mux := http.NewServeMux()
	mux.Handle("/list", http.HandlerFunc(db.List))
	mux.Handle("/price", http.HandlerFunc(db.Price))

	log.Fatal(http.ListenAndServe("localhost:3000", mux))
}
