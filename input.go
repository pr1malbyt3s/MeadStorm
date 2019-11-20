package main

import (
	"fmt"
	"net/http"
	"html/template"
	"html"
)

// InputPage function that serves a test page to display and handle an input field.
func InputPage(w http.ResponseWriter, r *http.Request) {
        // Prints the method used to access the page.
	fmt.Println("input request method:", r.Method)
        // Checks if the current request method is GET.
        // If so, parses and serves an HTML template.   
        if r.Method == "GET" {
                t, _ := template.ParseFiles("./assets/input.html")
                t.Execute(w, nil)
        // Checks if the request method is POST.
	// If so, records the input to the server, and prints the escaped output to the site user.
	} else if r.Method == "POST" {
	        fmt.Println("input data:", r.FormValue("data"))
                fmt.Fprintf(w, html.EscapeString(r.FormValue("data")))
        // If request method is not GET or POST, prints a method not allowed message.
	} else {
                fmt.Fprintf(w, "Method not allowed")
        }
}

