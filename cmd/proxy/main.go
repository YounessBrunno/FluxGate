package main

import (
	"log"
	"net/http"
	"net/url"

	proxy "github.com/YounessBrunno/FluxGate/internal/proxy"
)

func main() {
    // Parse the target URL for the reverse proxy ;)
    target, err := url.Parse("http://localhost:9000")
	if err != nil {
		log.Fatal("Failed to parse target URL:", err)
	}
    // Create a new proxy instance	
	p, err := proxy.NewProxy(target)
	if err != nil {
		log.Fatal("Failed to create proxy:", err)
	}

	log.Printf("Proxy listening on %s", ":8080")
	
	log.Fatal(http.ListenAndServe(":8080", p))
}
