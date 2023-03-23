package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"encoding/json"
)

var (
	config *ConfigurationObject
)

func main() {
	config = getConfig()

	if len(config.mappings) == 0 {
		log.Fatalln("No mapping found in configuration!")
		os.Exit(500)
	}

	if len(config.mappings) > 0 {
		startServe()
	}
}

func startServe() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", redirectToTarget)

	http.ListenAndServe(":80", mux)
}

func redirectToTarget(w http.ResponseWriter, r *http.Request) {
	log.Println("Request received for '", r.URL, "'")

	// TODO: get the target URL from the config
	targetUrl := "http://google.com"

	log.Println("Redirecting to '", targetUrl, "'")
	http.Redirect(w, r, targetUrl, http.StatusSeeOther)
}

func getConfig() *ConfigurationObject {
	// Get current directory
	directoryPath, directoryError := os.Getwd()
	if directoryError != nil {
		log.Fatalln("No directory found!")
		os.Exit(500)
	}

	// Load configuration file
	file, err := ioutil.ReadFile(directoryPath + "/config.json")
	if err != nil {
		log.Fatalln("No configuration file found")
		os.Exit(500)
	}

	data := &ConfigurationObject{}
	unmarshallError := json.Unmarshal(file, data)
	if unmarshallError != nil {
		log.Fatalln(unmarshallError)
		os.Exit(500)
	}

	log.Println(data)

	return data
}

type ConfigurationObject struct {
	test string `json:"source"`
	mappings []MappingObject `json:"mappings"`
}

type MappingObject struct {
	source string `json:"source"`
	target string `json:"target"`
}

type Pota struct {
	test string `json:"test"`
}