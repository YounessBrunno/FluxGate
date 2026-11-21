package main

import (
	"log"
	"net/http"

	_ "github.com/YounessBrunno/FluxGate/internal/proxy"
)

func main() {
	log.Printf("Proxy listening on %s", ":8080")
	http.ListenAndServe(":8080", nil)
}
