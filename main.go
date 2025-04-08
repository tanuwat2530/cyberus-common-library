package main

import (
	"fmt"
	"net/http"
)

func main() {

	// Start the server on port 8080
	fmt.Println("Starting server on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}

}
