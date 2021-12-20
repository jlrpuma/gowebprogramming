package main

import "testing"

func TestDecode(t *testing.T) {
	post, err := decode("post.json")
	if err != nil {
		t.Error()
	}

	if post.Id != 1 {
		t.Error("Decoded Post id is different", post.Id)
	}

	if post.Content != "Hello World!" {
		t.Error("Decoded Post content is different", post.Id)
	}
}
