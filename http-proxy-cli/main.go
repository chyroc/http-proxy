package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/chyroc/http-proxy"
)

func newConfig(usernameConfig, passwordConfig string) *http_proxy.Config {
	if usernameConfig == "" && passwordConfig == "" {
		return &http_proxy.Config{}
	}

	return &http_proxy.Config{
		AuthChecker: func(username, password string) error {
			if username == usernameConfig && password == passwordConfig {
				return nil
			}
			return fmt.Errorf("auth failed")
		},
	}
}

func main() {
	usernameConfig := os.Getenv("HTTP_PROXY_USERNAME")
	passwordConfig := os.Getenv("HTTP_PROXY_PASSWORD")
	serverHost := os.Getenv("HTTP_PROXY_SERVER_HOST")
	if serverHost == "" {
		serverHost = "127.0.0.1:8080"
	}
	proxy := http_proxy.NewHTTPProxy(newConfig(usernameConfig, passwordConfig))

	fmt.Printf("http proxy server start at %s\n", serverHost)
	if err := http.ListenAndServe(serverHost, proxy); err != nil {
		log.Fatalln(err)
	}
}
