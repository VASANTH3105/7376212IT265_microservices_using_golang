package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})

	http.HandleFunc("/name", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "VASANTH S")
	})

	http.HandleFunc("/regno", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "7376212IT265")
	})

	http.HandleFunc("/department", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Information Technology")
	})

	http.HandleFunc("/color", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "black")
	})

	fmt.Printf("Server running (port=8080), route: http://localhost:8080/helloworld\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)

	}

}
