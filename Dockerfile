FROM golang:alpine as builder

RUN apk add --no-cache git gcc libc-dev
RUN go get github.com/Kong/go-pluginserver

RUN mkdir /go-plugins
COPY go-plugins /go-plugins
RUN go build -buildmode plugin -o /go-plugins/go-key-checker.so /go-plugins/go-key-checker.go

FROM kong:2.0.1-alpine

COPY --from=builder /go/bin/go-pluginserver /usr/local/bin/
RUN mkdir /tmp/go-plugins
COPY --from=builder /go-plugins/go-key-checker.so /tmp/go-plugins/go-key-checker.so

COPY config.yml /etc/kong/config.yml
COPY kong.conf /etc/kong/kong.conf
