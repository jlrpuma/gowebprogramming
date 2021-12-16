package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"

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

func session(w http.ResponseWriter, r *http.Request) (sess data.Session) {
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
