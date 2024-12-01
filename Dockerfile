FROM golang:1.23.1

WORKDIR /app

# Copy Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy all source files into the container
COPY . .

# Expose the application port
EXPOSE 8080

# Run the Go application
CMD ["go", "run", "examples/main.go"]
