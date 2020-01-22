package main

import (
	"net/http"
	"log"
)

// Handler function that handles client requests and provides server's response.
func handlers() {
	// Handler for home page.
	http.HandleFunc("/", Home)
	// Handler for contact page.
	http.HandleFunc("/contact", ContactPage)
    // Handler for login page.
    http.HandleFunc("/login", LoginPage)
	// Listening function that serves web content. Will update to TLS once web application is built.
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}

func main() {
	// Call the handlers function.
	handlers()
}
