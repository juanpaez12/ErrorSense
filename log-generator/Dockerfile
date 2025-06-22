FROM golang:1.24.0-alpine AS builder

WORKDIR /app/log-generator

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

COPY log-generator/go.mod log-generator/go.sum ./
RUN go mod download
COPY log-generator/. .

RUN go build -o /usr/local/bin/log-generator cmd/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /usr/local/bin/log-generator .
RUN chmod +x ./log-generator

# Asegúrate del nombre exacto de tu archivo de plantilla aquí
COPY log-generator/.env.template .env.template

CMD ["./log-generator"]