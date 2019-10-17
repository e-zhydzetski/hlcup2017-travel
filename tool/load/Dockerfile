# build stage
FROM golang:1.12.9-alpine AS builder
RUN apk update \
    && apk add --no-cache \
        git \
        ca-certificates \
        tzdata \
    && update-ca-certificates
WORKDIR /workspace
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN cd tool/load && GOOS=linux GOARCH=amd64 go build -ldflags "-w -s" -o load

# final stage
FROM alpine:3.9.4
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
WORKDIR /app
ENTRYPOINT ["./load"]
COPY --from=builder /workspace/tool/load/load ./load