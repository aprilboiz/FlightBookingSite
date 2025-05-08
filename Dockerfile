FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

# Copy the source code
COPY . .

# Create documentation using swag
# Ensure swag is installed
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN swag init

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Final stage: create a minimal image
FROM alpine:latest

# Set working directory
WORKDIR /app

# Install necessary runtime dependencies
RUN apk --no-cache add ca-certificates

# Copy the binary from the builder stage
COPY --from=builder /app/main .

# Copy any other necessary files (config files, static assets, etc.)
COPY --from=builder /app/internal ./internal
COPY --from=builder /app/pkg ./pkg
COPY --from=builder /app/docs ./docs

# Expose the port that the application will run on
EXPOSE 8080

# Command to run the application
CMD ["./main"]