FROM golang:1.23-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o stress-test ./cmd

FROM alpine:3.19

WORKDIR /app

COPY --from=builder /app/stress-test .

RUN mkdir -p /app/logs

ENTRYPOINT ["/app/stress-test"]
