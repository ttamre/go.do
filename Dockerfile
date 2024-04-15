# Use the official Golang image as the base image
FROM golang:1.22.0

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules manifests
COPY go.mod .
COPY go.sum .
COPY Makefile .

# Download the dependencies
RUN make deps

# Copy the rest of the application source code
COPY . .

# Build the Go application
RUN make build

# Expose the port on which the Go application will run
EXPOSE 5000

# Install Redis
RUN apt-get update && apt-get install -y redis-server

# Expose the port for Redis (optional)
EXPOSE 5001

# Start the Redis daemon and run binary
CMD redis-server --port 5001 --daemonize yes && ./bin/godo