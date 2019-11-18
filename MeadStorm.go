package main

import (
	"fmt"
	"net/http"
	"log"
	"html/template"
)

// homePage function that provides the first page users visit on the application.
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "MeadStorm Home")
}

// inputPage function that serves a test page to display and hangle an input field.
func inputPage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("input.html")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		fmt.Println("input:", r.Form["data"])
	}
}

// Handler function that handles client requests and provides server's response.
func handlers() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/input", inputPage)
	// Listening function that serves web content. Will update to TLS once web application is built.
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}

func main() {
	handlers() // Call the handlers function
}
