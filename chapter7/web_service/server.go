package main

import (
	"encoding/json"
	"net/http"
	"path"
	"strconv"
)

type Post struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

func main() {
	// Server creation
	server := http.Server{
		Addr: ":8080",
	}
	// handle post CRUD
	http.HandleFunc("/post/", handleRequest)
	server.ListenAndServe()
}

// handling every verb
func handleRequest(w http.ResponseWriter, r *http.Request) {
	var err error
	switch r.Method {
	case "GET":
		err = handleGet(w, r)
	case "POST":
		err = handlePost(w, r)
	case "PUT":
		err = handlePut(w, r)
	case "DELETE":
		err = handleDelete(w, r)
	}
	if err != nil {
		http.Error(w, "Error", http.StatusInternalServerError)
		return
	}
}

func handleGet(w http.ResponseWriter, r *http.Request) (err error) {
	// TODO: based on the Base function is like
	// we are going to receive the number next to the
	// /post/1 <- second slash
	// then an Atoi is being used to convert the 1 string to number 1
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}
	// calling the method on data to get the info from the database
	post, err := retrieve(id)
	if err != nil {
		return
	}
	// Marshall info (passing from struc to json with identation)
	output, err := json.MarshalIndent(&post, "", "\t\t")
	if err != nil {
		return
	}
	// set ContentType to the ResponseWriter
	w.Header().Set("Content-Type", "application/json")
	// the Json outpiut from the Marshall is being added to the
	// response writer using the write method
	w.Write(output)
	return
}
