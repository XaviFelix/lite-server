package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/XaviFelix/lite-server.git/models"
)

// Root handler that handles requests to "/"
func RootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world")
}

// greet handler
func GreetHandler(w http.ResponseWriter, r *http.Request) {

	name := r.URL.Query().Get("name")
	var message string
	if name == "" {
		fmt.Fprintf(w, "Hello there, stranger!")
		message = "Hi there, stranger!"
	} else {
		fmt.Fprintf(w, "Hi, %s! welcome to my server!\n", name)
		message = fmt.Sprintf("Hi, %s! Welcome to my server!", name)
	}

	response := models.GreetResponse{Message: message}

	// Setting content type header
	w.Header().Set("Content-Type", "application/json")

	// Encode response to json
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
		return
	}
}
