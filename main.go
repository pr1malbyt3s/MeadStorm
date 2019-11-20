package main

import (
	"net/http"
	"log"
)

// Handler function that handles client requests and provides server's response.
func handlers() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/input", InputPage)
	// Listening function that serves web content. Will update to TLS once web application is built.
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}

func main() {
	handlers() // Call the handlers function
}
