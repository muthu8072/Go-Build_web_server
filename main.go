package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "parseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "post request sucessfull")
	name := r.FormValue("name")
	address := r.FormValue("address")
	email := r.FormValue("email")
	phno := r.FormValue("phno")
	qualification := r.FormValue("qualification")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
	fmt.Fprintf(w, "Email = %s\n", email)
	fmt.Fprintf(w, "Phno = %s\n", phno)
	fmt.Fprintf(w, "Qualification = %s\n", qualification)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "hello!")
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("starting server at port 8000\n")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
