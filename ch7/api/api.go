package api

import (
	"fmt"
	"net/http"
)

type dollars float64
type Database map[string]dollars

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

func (db Database) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/list":
		{
			for item, price := range db {
				fmt.Fprintf(w, "%s: %s\n", item, price)
			}
		}

	case "/price":
		{
			item := req.URL.Query().Get("item")
			price, ok := db[item]

			if !ok {
				w.WriteHeader(http.StatusNotFound)
				fmt.Fprintf(w, "item %s not found\n", item)
				return
			}

			fmt.Fprintf(w, "%s: %s\n", item, price)
		}
	default:
		{
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "page %s not found\n", req.URL)
		}
	}
}
