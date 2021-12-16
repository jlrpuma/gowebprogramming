package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Post struct {
	Id       int       `json:"id"`
	Content  string    `json:"content"`
	Author   Author    `json:"author"`
	Comments []Comment `json:"comments"`
}

type Author struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Comment struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

func main() {
	// getting the file
	jsonFile, err := os.Open("post.json")
	if err != nil {
		// just printing the error
		fmt.Println("Josn file cannot be opened", err)
		return
	}
	defer jsonFile.Close()

	// get file bytes
	jsonData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		// just printing the error
		fmt.Println("Josn file cannot be read", err)
		return
	}
	// make zero value useful
	var post Post
	// place the jsonData information on the direction of the post
	// that you've created above
	json.Unmarshal(jsonData, &post)
	fmt.Println(post)
}
