FROM golang:1.24-bullseye


WORKDIR /app

# Install build dependencies
RUN apt-get update && \
    apt-get install -y curl git cmake pkg-config libssl-dev && \
    rm -rf /var/lib/apt/lists/*

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .
RUN go mod tidy
# Build the application with the correct package path
RUN CGO_ENABLED=0 GOOS=linux go build -o api .

# Expose the API port
EXPOSE 8888

# Run the application
CMD ["./api"] 