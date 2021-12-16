package main

import (
	"encoding/json"
	"fmt"
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
	post := Post{
		Id:      1,
		Content: "Hello Go",
		Author: Author{
			Id:   1,
			Name: "Jose",
		},
		Comments: []Comment{
			{
				Id:      1,
				Content: "Hello Jose",
			},
			{
				Id:      2,
				Content: "hau?",
			}},
	}

	jsonFile, err := os.Create("post.json")
	if err != nil {
		fmt.Println("Error creating file", err)
		return
	}
	// Creating a New encoder
	encoder := json.NewEncoder(jsonFile)
	// encode the info that is stored on post
	// passing by reference
	err = encoder.Encode(&post)
	// handling err on Encoding
	if err != nil {
		fmt.Println("Error encoding", err)
		return
	}
}
