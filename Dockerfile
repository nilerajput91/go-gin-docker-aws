# Build Stage
FROM golang:1.20 AS builder

WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application
COPY . .

# Update Go modules
RUN go mod tidy

# Sync vendor directory if using vendor mode
RUN go mod vendor

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Final Stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates
WORKDIR /root/

# Copy the built binary from the builder stage
COPY --from=builder /app/main .

EXPOSE 8113

CMD ["./main"]
