package main

import (
	"net/http"

	"github.com/gowebprogramming/chapter2/data"
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
	// just a simple valiudation of the password that was queried
	// and the password received in the request (but encrypted)
	if user.Password == data.Encrypt(r.PostFormValue("password")) {
		session, err := user.CreateSession()
		if err != nil {
			danger(err, "Cannot create session")
		}

		// HttpOnly true allow Http or Https to access the cookie 
		// and no other non HTTP APIs like Javascript

		// expire attribute is not being setted because
		// this cookie will be a session cookie
		cookie := http.Cookie{
			Name:     "_cookie",
			Value:    session.Uuid,
			HttpOnly: true, 
		}
		// The cookie is being setted on the responseWriter
		http.SetCookie(w, &cookie)
		// Looks like this redirection cause a redirection on the browser to this page
		// this is interesting because you void the need of that redirection at frontend level
		// after a login is completer... (interesting)
		http.Redirect(w, r, "/", 302)
	} else {
		http.Redirect(w, r, "login", 302)
	}
}
