# Use the official Golang base image
FROM golang:latest AS builder

# Download dependencies
RUN go get github.com/prometheus/client_golang/prometheus
RUN go get github.com/prometheus/client_golang/prometheus/promhttp

# Set the working directory inside the container
WORKDIR /app

# Copy the application source code
COPY . .

# Build the Go application
RUN go build -o main -ldflags '-linkmode external -w -extldflags "-static"'

# Bundle Stage
FROM alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the executable from the builder stage
COPY --from=builder /app/main .

# Set the command to run the executable
CMD ["/app/main"]
