# Use official Golang image
FROM golang:1.24

# Set working directory
WORKDIR /app

# Copy Go modules and install dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application
COPY . .

# Build the Go application
RUN go build -o main .


# Expose port
EXPOSE 8080

# Run the application
CMD ["./main"]
