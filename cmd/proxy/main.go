package main

import (
	"log"
	"net/http"
	"net/url"
    "github.com/YounessBrunno/FluxGate/internal/config"
	"github.com/YounessBrunno/FluxGate/internal/proxy"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}

    // Parse the target URL for the reverse proxy ;)
    target, err := url.Parse("http://" + cfg.Host + ":" + cfg.ServerPort)
	if err != nil {
		log.Fatal("Failed to parse target URL:", err)
	}
    // Create a new proxy instance	
	p, err := proxy.NewProxy(target)
	if err != nil {
		log.Fatal("Failed to create proxy:", err)
	}

	log.Printf("Proxy listening on %s", ":" + cfg.ProxyPort)
	
	log.Fatal(http.ListenAndServe(":" + cfg.ProxyPort, p))
}
