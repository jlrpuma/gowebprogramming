package main

import (
	"fmt"
	"net/http"

	"github.com/gopatterns/github.com/gowebprogramming/chapter2/data"
)

// POST /authenticate
// authenticates the user given the email and password
// remember that this function is being passed as a handler
// to the multiplexer indicated route on the main file.
func authenticate(w http.ResponseWriter, r *http.Request) {
	// r.Form and r.PostForm gets initialized to get the information
	// from the POST request body
	err := r.ParseForm()
	// using the data layer of the application to get the user
	// based on the emial, once r.ParseForm populate
	// r.Form and r.PostForm the values can be reached
	// on this way r.PostFormValue({valueNeeded})
	user, err := data.UserByEmail(r.PostFormValue("email"))
	if err != nil {
		// usage of utility to log danger issues
		danger(err, "Cannot find user")
	}
	fmt.Println(user)
}
