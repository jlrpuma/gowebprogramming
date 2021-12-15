package main

import (
	"net/http"
)

func main() {

	// a new multiplexer gets created
	mux := http.NewServeMux()
	// FileServer provides a handler that will serve files from given directory
	files := http.FileServer(http.Dir(config.Static))

	// StripPrefix removes the _static_ part of the url and will
	// search for <application root>/{folderRequired}/{fileRequired}
	mux.Handle("/static/", http.StripPrefix("/static/", files))
}
