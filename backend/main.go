package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from the Go Backend!")
	})

	fmt.Println("Backend is running on port 3000")
	http.ListenAndServe(":3000", nil)
}