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

func handlePost(w http.ResponseWriter, r *http.Request) (err error) {
	len := r.ContentLength
	// creating an slice with an specified len
	body := make([]byte, len)
	// TODO: interesting way to get the body information
	r.Body.Read(body)

	// creating a post
	var post Post
	// to pass the Json information
	json.Unmarshal(body, &post)
	// post creation method is being executed
	err = post.create()
	if err != nil {
		return
	}
	// at this POST case the only information that the server transport
	// to the client is the 200 HTTP code, that says that everything goes ok
	// with tjhe request, howeverm a 201 could be used to right ?
	w.WriteHeader(200)
	return
}

func handlePut(w http.ResponseWriter, r *http.Request) (err error) {
	// get the resource id that wants to be updated
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}
	// retrieve the post by the id required
	post, err := retrieve(id)
	if err != nil {
		return
	}

	// extract the body content
	len := r.ContentLength
	// creating an slice with the size required for store the
	// body content
	body := make([]byte, len)
	// reading the Request Body and fill the body slice
	r.Body.Read(body)

	// TODO, this means that i can pass the id on the body
	// and i will be changing the id of the post, or even worst
	// i will be modifiying another post content...
	json.Unmarshal(body, &post)

	// using the update post method
	err = post.update()
	if err != nil {
		return
	}

	// The client just needs a confirmation that everything goes ok
	// so we put the status 200
	w.WriteHeader(200)
	return
}

func handleDelete(w http.ResponseWriter, r *http.Request) (err error) {
	// getting the resource id of the post that wants to be deleted
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}
	// TODO: is this really neded ?
	// Is needed because if the id doesn't exist 
	// the delete is not going to be successful
	post, err := retrieve(id)
	if err != nil {
		return
	}
	err = post.delete()
	if err != nil {
		return
	}
	w.WriteHeader(200)
	return
}
