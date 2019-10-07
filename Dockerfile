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
RUN GOOS=linux GOARCH=amd64 go build -ldflags "-w -s" -o app cmd/service/main.go

# final stage
FROM alpine:3.9.4
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
WORKDIR /app
ENTRYPOINT ["./app"]
ENV PORT 80
EXPOSE $PORT
ENV DOCKER 1
COPY --from=builder /workspace/app ./app