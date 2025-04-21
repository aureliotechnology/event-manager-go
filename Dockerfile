# Stage 1: Build
FROM golang:1.20.3-alpine AS builder
WORKDIR /app
RUN apk update && apk add --no-cache git
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o event-manager-go ./cmd/main.go

# Stage 2: Run
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/event-manager-go .
EXPOSE 8080
CMD ["./event-manager-go"]
