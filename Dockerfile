# Dockerfile for REST API
FROM golang:1.23-alpine

WORKDIR /app

# Copy go.mod and go.sum
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the application
RUN go build -o main .

# Expose the application port
#EXPOSE 8000

# Command to run the application
CMD ["./main"]