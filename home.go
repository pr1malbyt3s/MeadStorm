package main

import (
	"fmt"
	"net/http"
	"html/template"
	"log"
)

// Home function that handles requests to site root and serves a static home page.
func Home(w http.ResponseWriter, r *http.Request) {
	// Prints the request method used to access the home page.
	fmt.Println("home request method:", r.Method)
	t, err := template.ParseFiles("./assets/home.html")
	if err != nil {
		log.Fatal("Error parsing home.html: ", err)
	}
	t.Execute(w, nil)
}
