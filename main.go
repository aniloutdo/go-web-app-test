package main

import (
	"fmt"
//	"log"
	"net/http"
        "os"
)

func main() {
	// Register a handler function for the root path "/"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})

        http.HandleFunc("/secret", func(w http.ResponseWriter, r *http.Request) {
                fmt.Fprint(w, os.Getenv("API_SECRET"))
//                secret := os.Getenv("API_SECRET")
//                fmt.Fprintln(w, secret)
        })


	// Start the HTTP server on port 8080
        http.ListenAndServe(":8080", nil)
//	port := ":8080"
//	fmt.Printf("Server starting on port %s\n", port)
//	log.Fatal(http.ListenAndServe(port, nil))
}
