FROM golang:1.24.0-alpine AS builder

WORKDIR /app/anomaly-processor

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

COPY anomaly-processor/go.mod anomaly-processor/go.sum ./
COPY log-generator /app/log-generator
RUN go mod download
COPY anomaly-processor/. .

RUN go build -o /usr/local/bin/anomaly-processor cmd/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /usr/local/bin/anomaly-processor .
RUN chmod +x ./anomaly-processor

# Asegúrate del nombre exacto de tu archivo de plantilla aquí
COPY anomaly-processor/.env.template .env.template

CMD ["./anomaly-processor"]