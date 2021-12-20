package main

import (
	"encoding/json"
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

}

func decode(filename string) (post Post, err error) {
	file, err := os.Open(filename)
	if err != nil {
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&post)
	if err != nil {
		return
	}
	return
}

func unmarshal(filename string) (post Post, err error) {
	file, err := os.Open(filename)
	if err != nil {
		return
	}
	defer file.Close()

	jsonData, err := ioutil.ReadAll(file)
	if err != nil {
		return
	}

	err = json.Unmarshal(jsonData, &post)
	if err != nil {
		return
	}
	return
}
