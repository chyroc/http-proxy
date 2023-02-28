FROM golang:1.19 AS build

ENV GOPATH /go
WORKDIR /app

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /go/bin/http-proxy-cli ./http-proxy-cli/main.go

RUN strip /go/bin/http-proxy-cli
RUN test -e /go/bin/http-proxy-cli

FROM alpine:latest

LABEL org.opencontainers.image.source=https://github.com/chyroc/http-proxy
LABEL org.opencontainers.image.description="Simple Http Proxy Server Supporting Authentication."
LABEL org.opencontainers.image.licenses="Apache-2.0"

COPY --from=build /go/bin/http-proxy-cli /bin/http-proxy-cli

ENTRYPOINT ["/bin/http-proxy-cli"]
