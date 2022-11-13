package main

import (
	"log"
	"net/http"
)

func redirectToTarget(w http.ResponseWriter, r *http.Request) {
	log.Println("Request received for '", r.URL, "'")

	targetUrl := "http://google.com"

	log.Println("Redirecting to '", targetUrl, "'")
	http.Redirect(w, r, targetUrl, http.StatusSeeOther)
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", redirectToTarget)

	http.ListenAndServe(":80", mux)
}