package main

import (
	"fmt"
	"time"
	"log"
	"net/http"
	"encoding/json"
	"github.com/YounessBrunno/FluxGate/internal/config"
)

func main() {
    mux := http.NewServeMux()

    cfg, err := config.Load()
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}

	log.Printf("Backend listening on %s", ":" + cfg.ServerPort)

	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to the backend server!")
	})

	mux.HandleFunc("/slow", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(3 * time.Second)
		fmt.Fprintf(w, "This is a slow endpoint!")
	})
    
	type User struct {
		Name string `json:"name"`
	}

	users := []User{
		{Name: "Alice"},
		{Name: "Bob"},
		{Name: "Charlie"}, 
    }

	mux.HandleFunc("GET /users", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(users)
	})
    
	mux.HandleFunc("POST /users", func(w http.ResponseWriter, r *http.Request) {
		var user User

		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		if user.Name == "" {
			http.Error(w, "Missing name", http.StatusBadRequest)
			return
		}
         
		users = append(users, user)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)

		json.NewEncoder(w).Encode(user)
    }) 

    server := &http.Server{
		Addr:    ":" + cfg.ServerPort,
		Handler: mux,
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
