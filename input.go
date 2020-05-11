package main

import (
	"fmt"
	"net/http"
	"html/template"
	//"html"
)

// ContactPage function that mimicks a contact form.
func ContactPage(w http.ResponseWriter, r *http.Request) {
        // Prints the method used to access the page.
	fmt.Println("contact request method:", r.Method)
        // Checks if the current request method is GET.
        // If so, parses and serves an HTML template.   
        if r.Method == "GET" {
                t, _ := template.ParseFiles("./assets/contact.html")
                t.Execute(w, nil)
        // Checks if the request method is POST.
        // If so, records the input to the server, and sends input data to MongoDB.
	} else if r.Method == "POST" {
            fmt.Fprintf(w, "Thanks for your message!")
	        fmt.Println("contact name data:", r.FormValue("name"))
                //fmt.Fprintf(w, html.EscapeString(r.FormValue("name")))
            fmt.Println("contact email data:", r.FormValue("email"))
                //fmt.Fprintf(w, html.EscapeString(r.FormValue("email")))
            fmt.Println("contact message data:", r.FormValue("message"))
                //fmt.Fprintf(w, html.EscapeString(r.FormValue("message")))
            //MongoSendContact(r.Method, r.FormValue("name"), r.FormValue("email"), r.FormValue("message"))
        // If request method is not GET or POST, prints a method not allowed message.
	} else {
                fmt.Fprintf(w, "Method not allowed")
        }
}

// LoginPage function that mimicks a login portal.
func LoginPage(w http.ResponseWriter, r *http.Request) {
        // Prints the method used to access the page.
	fmt.Println("login request method:", r.Method)
        // Checks if the current request method is GET.
        // If so, parses and serves an HTML template.   
        if r.Method == "GET" {
                t, _ := template.ParseFiles("./assets/portal.html")
                t.Execute(w, nil)
        // Checks if the request method is POST.
        // If so, records the input to the server, and sends input data to MongoDB.
	} else if r.Method == "POST" {
            fmt.Fprintf(w, "Login incorrect. Try again.")
	        fmt.Println("login username data:", r.FormValue("username"))
                //fmt.Fprintf(w, html.EscapeString(r.FormValue("username")))
            fmt.Println("login password data:", r.FormValue("password"))
                //fmt.Fprintf(w, html.EscapeString(r.FormValue("password")))
            //MongoSendLogin(r.Method, r.FormValue("username"), r.FormValue("password"))
        // If request method is not GET or POST, prints a method not allowed message.
	} else {
                fmt.Fprintf(w, "Method not allowed")
        }
}
