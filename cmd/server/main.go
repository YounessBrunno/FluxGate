package main

import (
	"fmt"
	"time"
	"log"
	"net/http"
)

func main() {
	log.Printf("Backend listening on %s", ":9000")

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "path: %s", r.URL.Path)
	})

	http.HandleFunc("/slow", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(3 * time.Second)
		fmt.Fprintf(w, "This is a slow endpoint!")
	})

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		users := []string{"Alice", "Bob", "Charlie"}
		for _, user := range users {
			fmt.Fprintf(w, "User: %s\n", user)
		}
	})

	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		log.Fatal("Failed to start backend:", err)
	}
}
