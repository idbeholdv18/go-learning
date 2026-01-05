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
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}
