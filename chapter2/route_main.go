package main

import (
	"net/http"

	"github.com/gowebprogramming/chapter2/data"
)

func index(w http.ResponseWriter, r *http.Request) {
	threads, err := data.Threads()
	if err != nil {
		error_message(w, r, "Cannot get threads")
	} else {
		_, err := session(w, r)
		if err != nil {
			generateHtml(w, threads, "layout", "public.navbar", "index")
		} else {
			generateHtml(w, threads, "layout", "public.navbar", "index")
		}
	}
}
