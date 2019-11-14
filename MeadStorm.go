package main

import (
	"fmt"
	"net/http"
)

func main() {
	handle()
	listen()
}

func handle() {
	// Handler function that handles client requests and provides server's response.
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "MeadStorm Test")
	})
}

func listen() {
	// Listening function that serves web content on 443. Will update to TLS once web application is built.
	http.ListenAndServe(":8080", nil)
}
