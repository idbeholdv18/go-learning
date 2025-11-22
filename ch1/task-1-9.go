package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)

		if err != nil {
			fmt.Fprintf(os.Stderr, "get err: %v", err)
			os.Exit(1)
		}

		fmt.Printf("%s: %s\n", url, resp.Status)

		resp.Body.Close()
	}
}
