package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	log.Printf("Backend listening on %s", ":9000")

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})

	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		log.Fatal("Failed to start backend:", err)
	}
}
