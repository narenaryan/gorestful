package main

import (
	"fmt"
	"net/http"
        "math/rand"
)

type CustomServeMux struct {
}

func (p *CustomServeMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		sayhelloName(w, r)
		return
	}
	http.NotFound(w, r)
	return
}

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,"Your random number is: %f", rand.Float64())
}

func main() {
	mux := &CustomServeMux{}
	http.ListenAndServe(":8000", mux)
}
