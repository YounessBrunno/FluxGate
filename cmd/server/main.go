package main

import (
	"fmt"
	"time"
	"log"
	"net/http"
	"github.com/YounessBrunno/FluxGate/internal/config"
)

func main() {

    cfg, err := config.Load()
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}

	log.Printf("Backend listening on %s", ":" + cfg.ServerPort)

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
    
    server := &http.Server{
		Addr:    ":" + cfg.ServerPort,
		Handler: nil,
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      15 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	err = server.ListenAndServe()
	if err != nil {
		log.Fatal("Failed to start backend:", err)
	}
}
