ARG GOLANG_VERSION="1.25"
ARG ALPINE_VERSION="3.22"

FROM golang:${GOLANG_VERSION}-alpine${ALPINE_VERSION} AS builder

RUN apk add --no-cache ca-certificates

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /tmp/go-expo ./cmd/go-expo

FROM alpine:${ALPINE_VERSION}

RUN addgroup -S 1111 && adduser -S 1111 -G 1111
RUN mkdir /etc/go-expo && chown -R 1111:1111 /etc/go-expo

COPY --from=builder /tmp/go-expo /usr/local/bin/go-expo

USER 1111

CMD ["/usr/local/bin/go-expo", "--conf", "/etc/go-expo/config.yaml"]