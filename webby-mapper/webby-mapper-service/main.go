package main

import (
	"log"
	"net/http"

	viper "github.com/spf13/viper"
)

var (
	config *ConfigurationObject
)

func main() {
	// Get the configuration file
	viper.SetConfigFile("./webby-mapper-service/config.json")

	// Read the configuration file
	err := viper.ReadInConfig()

	// Handle errors reading the config file
	if err != nil {
		errorMessage := "Error while reading the config file"
		log.Print(errorMessage, err);
		panic(errorMessage)
	}

	config := &ConfigurationObject{}
	// Unmarshal the config file
	unmarshalErr := viper.Unmarshal(config)
	// Handle errors reading the config file
	if unmarshalErr != nil {
		errorMessage := "Error while Unmarshalling the config file"
		log.Print(errorMessage, unmarshalErr);
		panic(errorMessage)
	}

	mappings := viper.Get("mappings")
	
	// TODO: set the array on the config prop
	//config.mappings = make([]MappingObject, mappings.len())

	log.Println(mappings)
	log.Println(config)
	log.Println(config.mappings)

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

type ConfigurationObject struct {
	mappings []MappingObject
}

type MappingObject struct {
	source string
	target string
}