package main

import (
	"go-learning/ch7/api"
	"log"
	"net/http"
)

func main() {
	db := api.Database{"shoes": 30, "socks": 5}

	log.Fatal(http.ListenAndServe("localhost:3000", db))
}
