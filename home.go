package main

import (
	"net/http"
	"html/template"
	"log"
)

func Home(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./assets/home.html")
	if err != nil {
		log.Fatal("Error parsing home.html: ", err)
	}
	t.Execute(w, nil)
}
