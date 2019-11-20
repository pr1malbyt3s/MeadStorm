package main

import (
	"fmt"
	"net/http"
	"html/template"
	"html"
)

// InputPage function that serves a test page to display and handle an input field.
func InputPage(w http.ResponseWriter, r *http.Request) {
        fmt.Println("method:", r.Method) // Prints the current request method.
        // Checks if the current request method is GET.
        // If so, parses and serves an HTML template.   
        if r.Method == "GET" {
                t, _ := template.ParseFiles("./assets/input.html")
                t.Execute(w, nil)
        } else if r.Method == "POST" {
	        fmt.Println("input:", r.FormValue("data"))
                fmt.Fprintf(w, html.EscapeString(r.FormValue("data")))
        } else {
                fmt.Fprintf(w, "Method not allowed")
        }
}

