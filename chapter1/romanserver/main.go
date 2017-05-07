package main

import (
	"fmt"
	"html"
	"net/http"
	"time"
        "data/romanNumerals"
)

func main() {
	// http package has methods for dealing with requests
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
                fmt.Println(r.URL)
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})
	// Create a server and run it on 8000 port
	s := &http.Server{
		Addr:           ":8000",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
