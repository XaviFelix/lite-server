package main

import (
	"fmt"
	"net/http"
)

// Root handler that handles requests to "/"
func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world")
}

// greet handler
func greetHandler(w http.ResponseWriter, r *http.Request) {

	name := r.URL.Query().Get("name")
	if name == "" {
		fmt.Fprintf(w, "Hello there, stranger!")
	} else {
		fmt.Fprintf(w, "Hi, %s! welcome to my server!", name)
	}
}

func main() {

	// registering my handlers:
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/greet", greetHandler)

	// begin the server on port 8080
	fmt.Println("Server starting on :8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server: ", err)
	}

}
