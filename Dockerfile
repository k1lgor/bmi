# Start from the official Go image for building
FROM golang:1.23-alpine AS builder

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the BMI binary
RUN go build -o bmi ./cmd/bmi

# Start a minimal image for running
FROM alpine:latest

WORKDIR /app

# Copy the binary from the builder
COPY --from=builder /app/bmi .

# Expose the web port (if using web mode)
EXPOSE 8080

# Default command (can be overridden)
ENTRYPOINT ["./bmi"]
