package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

	jsonContent, err := json.Marshal(post)
	if err != nil {
		fmt.Println("Error on marshall", err)
	}
	err = ioutil.WriteFile("post.json", jsonContent, 0644)
	if err != nil {
		fmt.Println()
	}

}
