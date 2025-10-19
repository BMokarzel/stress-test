# Etapa 1: Build
FROM golang:1.23-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o stress-test ./cmd

# Etapa 2: Runtime
FROM alpine:3.19

WORKDIR /app

# Copia apenas o binário
COPY --from=builder /app/stress-test .

# Cria o diretório de logs (fora do binário)
RUN mkdir -p /app/logs

ENTRYPOINT ["/app/stress-test"]
