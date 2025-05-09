# Use Golang base image
FROM golang:1.21

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum
COPY go.mod ./
COPY go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application
COPY . .

# Build the application
RUN go build -o app .

# Expose port (update if your app uses a different port)
EXPOSE 8080

# Run the app
CMD ["./app"]
