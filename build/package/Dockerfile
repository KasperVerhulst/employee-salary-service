# Build stage
FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o main .

# Final stage
FROM debian:trixie-slim
MAINTAINER Kasper Verhulst

# Add a non-root user and group
RUN groupadd -r appuser && useradd --no-log-init -r -g appuser appuser

# Copy binary and set permissions
COPY --from=builder /app/main .
RUN chown appuser:appuser /main
USER appuser

EXPOSE 8080
CMD ["./main"]