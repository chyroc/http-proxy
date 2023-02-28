package http_proxy

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"
)

func getAuth(request *http.Request) (string, string, error) {
	authorization := request.Header.Get("Proxy-Authorization")
	if authorization == "" {
		return "", "", nil
	}
	if !strings.HasPrefix(authorization, "Basic ") {
		return "", "", fmt.Errorf("invalid authorization")
	}
	d := authorization[6:]
	bs, err := base64.StdEncoding.DecodeString(d)
	if err != nil {
		return "", "", fmt.Errorf("invalid authorization")
	}
	dd := strings.SplitN(string(bs), ":", 2)
	if len(dd) != 2 {
		return "", "", fmt.Errorf("invalid authorization")
	}
	username := dd[0]
	password := dd[1]
	if username == "(nil)" {
		username = ""
	}
	if password == "(nil)" {
		password = ""
	}
	return username, password, nil
}
