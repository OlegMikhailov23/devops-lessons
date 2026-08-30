## --- Stage 1: Build ---
FROM golang:1.19 AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o momo-store ./cmd/api

## --- Stage 2: Run ---
FROM debian:bullseye-slim
RUN useradd -m appuser
USER appuser

WORKDIR /app
COPY --from=builder app/momo-store .

EXPOSE 8081

ENTRYPOINT [ "./momo-store" ]
