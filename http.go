package http_proxy

import (
	"io"
	"net/http"
)

func httpProxy(writer http.ResponseWriter, request *http.Request) {
	if request.Body != nil {
		defer request.Body.Close()
	}

	proxyRequest, err := http.NewRequest(request.Method, request.URL.String(), request.Body)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	proxyRequest.Header = request.Header
	proxyRequest.Header.Set("X-Forwarded-For", request.Header.Get("User-Agent"))
	proxyRequest.Header.Set("X-Forwarded-Host", request.Host)
	proxyRequest.Header.Set("X-Forwarded-Proto", request.Proto)

	proxyResponse, err := http.DefaultClient.Do(proxyRequest)
	if proxyResponse != nil && proxyResponse.Body != nil {
		defer proxyResponse.Body.Close()
	}
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadGateway)
		return
	}

	writer.WriteHeader(proxyResponse.StatusCode)
	for k, v := range proxyResponse.Header {
		for _, vv := range v {
			writer.Header().Add(k, vv)
		}
	}
	io.Copy(writer, proxyResponse.Body)
}
