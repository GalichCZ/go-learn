package main

import (
	"fmt"
	"log"
	"net/http"
)

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/about" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method is not supported", http.StatusNotFound)
		return
	}

	http.ServeFile(w, r, "./static/about.html")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
	}
	fmt.Fprintf(w, "POST request success\n")

	name := r.FormValue("name")
	surname := r.FormValue("surname")

	fmt.Fprintf(w, "Name: %s \nSurname: %s", name, surname)
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/form", formHandler)

	fmt.Println("Starting server on 8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
