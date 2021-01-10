package main

import (
	"errors"
	"fmt"
	"html"
	"log"
	"net/http"
)

//Hello generates greet string
func Hello(name string) (string, error) {
	if name == "" {
		return name, errors.New("Name is empty value")
	}
	return "Hi " + name, nil
}

func main() {
	m, err := Hello("jD")
	if err != nil {
		log.Fatalf("Error getting greet")
	}
	fmt.Println(m)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
	fmt.Println("End")
}
