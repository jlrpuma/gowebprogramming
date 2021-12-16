package main

import (
	"net/http"

	"github.com/gowebprogramming/chapter2/data"
)

func index(w http.ResponseWriter, r *http.Request) {
	_, err := data.Threads()
	if err != nil {
		error_message(w, r, "Cannot get threads")
	} else {
		_, err := session(w, r)
		if err != nil {

		} else {

		}
	}
}
