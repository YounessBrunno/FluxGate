package main

import (
	"fmt"
	"net/http"
	"time"
)


func main() {
   http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
   })

   http.HandleFunc("/slow", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(3 * time.Second)
		fmt.Fprintf(w, "Slept for 3 seconds!")
   })

   http.ListenAndServe(":9000", nil)
}
