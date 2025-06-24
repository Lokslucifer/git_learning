# Use Go base image
FROM golang:1.23.4-alpine3.19

# Set working directory
WORKDIR /app

# Copy files
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

# Build the Go app
RUN go build -o main ./cmd

# Run the app
CMD ["./main"]
