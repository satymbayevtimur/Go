package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		_, err := fmt.Fprintf(w, "ParseForm() err: %v", err)

		if err != nil {
			return
		}
	}

	_, err := fmt.Fprintf(w, "POST request successful\n")

	if err != nil {
		return
	}

	name := r.Form.Get("name")
	address := r.Form.Get("address")

	_, err = fmt.Fprintf(w, "Name = %s\nAddress = %s\n", name, address)

	if err != nil {
		return
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.NotFound(w, r)

		return
	}

	if r.Method != "GET" {
		http.NotFound(w, r)
	}

	_, err := fmt.Fprintf(w, "Hello!")

	if err != nil {
		return
	}
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))

	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Listening on port 8080\n")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
