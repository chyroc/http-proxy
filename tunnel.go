package http_proxy

import (
	"io"
	"net"
	"net/http"
	"time"
)

func tcpTunnel(w http.ResponseWriter, host string) {
	if len(host) > 2 && host[0:2] == "//" {
		host = host[2:]
	}
	destConn, err := net.DialTimeout("tcp", host, 10*time.Second)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}

	w.WriteHeader(http.StatusOK)

	hijacker, ok := w.(http.Hijacker)
	if !ok {
		http.Error(w, "Hijacker not supported", http.StatusInternalServerError)
		return
	}

	clientConn, _, err := hijacker.Hijack()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}

	go copyCloser(destConn, clientConn)
	go copyCloser(clientConn, destConn)
}

func copyCloser(destination io.WriteCloser, source io.ReadCloser) {
	defer destination.Close()
	defer source.Close()

	io.Copy(destination, source)
}
