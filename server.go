package http_proxy

import (
	"fmt"
	"net/http"
)

type proxy struct {
	config *Config
}

type Config struct {
	AuthChecker func(username, password string) error
}

func NewHTTPProxy(config *Config) http.Handler {
	return &proxy{config: config}
}

func (r *proxy) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	url := request.URL.String()
	fmt.Printf("[proxy] method='%s', url='%s'\n", request.Method, url)

	username, password, err := getAuth(request)
	if r.config.AuthChecker != nil {
		if err != nil {
			http.Error(writer, err.Error(), http.StatusProxyAuthRequired)
			return
		}
		if err := r.config.AuthChecker(username, password); err != nil {
			http.Error(writer, err.Error(), http.StatusProxyAuthRequired)
			return
		}
	}

	if request.Method == http.MethodConnect {
		tcpTunnel(writer, url)
		return
	}

	httpProxy(writer, request)
}
