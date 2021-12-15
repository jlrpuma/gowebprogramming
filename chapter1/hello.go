package main

import (
	"fmt"
	"net/http"
)

// ResponseWriter interface
// Request pointer struct
// at this case we take some information from the request to send it through the
// response writer\
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world %s!", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
