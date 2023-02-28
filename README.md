# http-proxy

Simple Http Proxy Server Supporting Authentication.

## Usage

### By Docker

```shell
docker run -d \
  -p 8080:8080 \
  -e HTTP_PROXY_SERVER_HOST=127.0.0.1:8080 \
  -e HTTP_PROXY_USERNAME=foo \
  -e HTTP_PROXY_PASSWORD=bar \
  ghcr.io/chyroc/http-proxy-cli:0.1.0
```

### By Binary

```shell
go install github.com/chyroc/http-proxy/http-proxy-cli@latest

export HTTP_PROXY_SERVER_HOST=127.0.0.1:8080
export HTTP_PROXY_USERNAME=foo
export HTTP_PROXY_PASSWORD=bar

http-proxy-cli
```
