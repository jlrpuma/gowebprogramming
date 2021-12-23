package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestMain(m *testing.M) {
	setUp()
}

// creating a mux to pass the right handler
// that is located on server.go
var mux *http.ServeMux

// response recoder
var writer *httptest.ResponseRecorder

func setUp() {
	mux := http.NewServeMux()
	mux.HandleFunc("/post/", requestHandler)
}

func TestHandleGet(t *testing.T) {
	request, _ := http.NewRequest("GET", "/post/1", nil)
	// once the request is done, writer will have the response value
	mux.ServeHTTP(writer, request)

	// validating the HTTP code that is being stored on the writer
	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}
	var post Post
	// unmarshaling the Body on the writer
	json.Unmarshal(writer.Body.Bytes(), &post)
	if post.Id != 1 {
		t.Error("Cannot retrieve JSON post")
	}
}

func TestHandlePost(t *testing.T) {
	// the json String is being embbeded on a Reader provided by the stirngs package
	json := strings.NewReader(`{"content":"Contenido del post", "author":"Jose"}`)
	request, _ := http.NewRequest("POST", "/post/", json)
	mux.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}
}
