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
	server := &http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/post/", requestHandler)
	server.ListenAndServe()
}

func requestHandler(w http.ResponseWriter, r *http.Request) {
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
		// w.WriteHeader(500)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func handleGet(w http.ResponseWriter, r *http.Request) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}
	var post Post
	post, err = retrieve(id)
	if err != nil {
		return
	}

	output, err := json.Marshal(&post)
	if err != nil {
		return
	}

	w.Header().Set("Content-Type:", "application/json")
	w.Write(output)
	return
}

func handlePost(w http.ResponseWriter, r *http.Request) (err error) {
	len := r.ContentLength
	content := make([]byte, len)
	r.Body.Read(content)
	var post Post
	json.Unmarshal(content, &post)
	err = post.create()
	if err != nil {
		return
	}
	w.WriteHeader(201)
	return
}

// /post/1
func handleDelete(w http.ResponseWriter, r *http.Request) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}
	var post Post
	post, err = retrieve(id)
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

// /post/ handle request body
func handlePut(w http.ResponseWriter, r *http.Request) (err error) {
	len := r.ContentLength
	content := make([]byte, len)
	r.Body.Read(content)
	var post Post
	json.Unmarshal(content, &post)
	err = post.update()
	if err != nil {
		return
	}
	w.WriteHeader(200)
	return
}
