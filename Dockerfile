FROM golang:1.23-alpine AS builder
WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o microservice ./cmd/server

FROM alpine:3.21
WORKDIR /app
RUN addgroup -S app && adduser -S app -G app

COPY --from=builder /app/microservice ./microservice

USER app
EXPOSE 8080
ENTRYPOINT ["/app/microservice"]
