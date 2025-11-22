package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		final_url := url

		if !strings.HasPrefix(url, "http://") {
			final_url = "http://" + url
		}

		resp, err := http.Get(final_url)

		if err != nil {
			fmt.Fprintf(os.Stderr, "get error: %v\n", err)
			os.Exit(1)
		}

		_, err = io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()

		if err != nil {
			fmt.Fprintf(os.Stderr, "read error: %s: %v\n", final_url, err)
			os.Exit(1)
		}
	}
}
