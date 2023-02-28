package http_proxy

import (
	"fmt"
	"net/http"
)

func Server(writer http.ResponseWriter, request *http.Request) {
	url := request.URL.String()
	fmt.Printf("[proxy] method='%s', url='%s'\n", request.Method, url)

	if request.Method == http.MethodConnect {
		tcpTunnel(writer, url)
		return
	}

	httpProxy(writer, request)
}
