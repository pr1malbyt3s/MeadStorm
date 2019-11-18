package main

import (
	"fmt"
	"net/http"
	"log"
	"html/template"
)

/*// homePage function that provides the first page users visit on the application.
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "MeadStorm Home")
}
*/
// InputPage function that serves a test page to display and handle an input field.
func InputPage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) // Prints the current request method.
	// Checks if the current request method is GET.
	// If so, parses and serves an HTML template. 	
	if r.Method == "GET" {
		t, err := template.ParseFiles("./assets/input.html")
		if err != nil {
			log.Fatal("ParseFiles: ", err)
		}
		t.Execute(w, nil)
	} else if r.Method == "POST" {
		fmt.Println("input:", r.FormValue("data"))
		fmt.Fprintf(w, r.FormValue("data"))
	} else {
		fmt.Fprintf(w, "Method not allowed")
	}
}

// Handler function that handles client requests and provides server's response.
func handlers() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/InputPage", InputPage)
	// Listening function that serves web content. Will update to TLS once web application is built.
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}

func main() {
	handlers() // Call the handlers function
}
