package config


import (
	"github.com/joho/godotenv"
	"os"
	"strings"
)


type Config struct {
 	ServerPort string
	ProxyPort string
	Host string
}


func Load() (Config, error) {

	_ = godotenv.Load(".env.local", ".env")

    serverPort := stringFromEnv("SERVER_PORT", "9000")
	proxyPort := stringFromEnv("PROXY_PORT", "8080")
	host := stringFromEnv("HOST", "localhost")

	return Config{
		ServerPort: serverPort,
		ProxyPort: proxyPort,
		Host: host,
	}, nil
}

func stringFromEnv(key string, fallback string) string {
	value := strings.TrimSpace(os.Getenv(key))
	if value == "" {
		return fallback
	}

	return value
}
