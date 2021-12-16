package main

import (
	"encoding/json"
	"fmt"
	"io"
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
	jsonFile, err := os.Open("post.json")
	if err != nil {
		fmt.Println("Error opening file", err)
	}
	defer jsonFile.Close()

	// a Decoder is created based on the file
	decoder := json.NewDecoder(jsonFile)
	for {
		var post Post
		// pass to the decoder decode method the post reference
		err := decoder.Decode(&post)
		//handle the EOF returned by the Decode
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error on decoding", err)
		}
		fmt.Println(post)
	}
}
