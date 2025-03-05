package main

import (
	"fmt"
	"net/http"

	"github.com/XaviFelix/lite-server.git/handlers"
)

func main() {

	// registering my handlers:
	http.HandleFunc("/", handlers.RootHandler)
	http.HandleFunc("/greet", handlers.GreetHandler)

	// begin the server on port 8080
	fmt.Println("Server starting on :8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server: ", err)
	}
}
