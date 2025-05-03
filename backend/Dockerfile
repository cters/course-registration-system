# STEP 1: Build the Go app
FROM golang:1.24 AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main cmd/server/main.go

# STEP 2: Create a minimal image to run the Go app
FROM alpine:latest
WORKDIR /app

COPY --from=builder /app/main .

# Expose the port the app runs on
EXPOSE 8000

# Command to run the executable
CMD ["./main"]