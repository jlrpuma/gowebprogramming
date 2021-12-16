package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"

	"github.com/gowebprogramming/chapter2/data"
)

type Configuration struct {
	Address      string
	ReadTimeout  int64
	WriteTimeout int64
	Static       string
}

// variable config will be accesible from main because they are on the same package
var config Configuration

var logger *log.Logger

// init function loads the configuration file into a config variable declared above
func init() {
	loadConfig()
}

func loadConfig() {
	// loading configuration JSON file
	file, err := os.Open("config.json")
	// handling error reported by os.Open
	if err != nil {
		log.Fatal("Cannot open config file", err)
	}

	// A decoder is created to handle de JSON on the file
	decoder := json.NewDecoder(file)
	// Give us a variable to store the information that will be decoded from the file
	// this variables does not require :=, because config will containt an Empty Configuration
	// struct.
	config = Configuration{}
	// decoder.Decode just returns an error at any case the Decoding process fails
	err = decoder.Decode(&config)
	// handle error reported by the Decode process
	if err != nil {
		log.Fatal("Cannot get the configuration from the file", err)
	}

}

func session(w http.ResponseWriter, r *http.Request) (sess data.Session, err error) {
	// we try to get the cookie that can be stored on the request
	cookie, err := r.Cookie("_cookie")
	// if the cookie is found
	if err == nil {
		// creating a session based on the Value that the cookie has stored
		// remember that value is the Uuid that we put on the authentication
		// handler
		sess = data.Session{Uuid: cookie.Value}
		//
		if ok, _ := sess.Check(); !ok {
			err = errors.New("Invalid session")
		}
	}
	return
}

// danger allows us to use the logger in a more centralized way
// using our utils file to log any classified danger issue.
func danger(args ...interface{}) {
	logger.SetPrefix("ERROR")
	logger.Println(args...)
}

// this utility was meant for redirect to the appropiate page
// the user based on an error, providing a message
func error_message(w http.ResponseWriter, r *http.Request, message string) {
	// TODO: don't really understand why the strings are concatenated in this way
	url := []string{"/err?msg=", message}
	http.Redirect(w, r, strings.Join(url, ""), 302)
}

func generateHtml(w http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, filename := range filenames {
		// appending the address of every file on the files slice
		files = append(files, fmt.Sprintf("templates/%s", filename))
	}

	// ParseFiles, will convert those plain files on templates
	// that will be validated by Must method
	templates := template.Must(template.ParseFiles(files...))
	// then the templates will be executed with the data provided
	// you cn see that the data can be opened to receive an interface
	// thats because the template can require a set of multiple kinds of data
	// is up to you if you send the correct one based on the filenames that
	// you send to this function
	templates.ExecuteTemplate(w, "layout", data)
}
