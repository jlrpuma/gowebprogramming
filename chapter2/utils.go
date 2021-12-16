package main

import (
	"encoding/json"
	"log"
	"os"
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
