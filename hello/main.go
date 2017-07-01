package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	msg := os.Getenv("MESSAGE")
	if msg == "" {
		msg = "`MESSAGE` not set..."
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, msg)
	})

	http.ListenAndServe(":9999", nil)
}
