# Stage 1: Build the Go application
FROM golang:1.23-alpine AS build

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if go.mod and go.sum are not changed
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go app
RUN go build -o main .

# Stage 2: Run the application
FROM alpine:latest

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the Pre-built binary from the build stage
COPY --from=build /app/main .

# Copy the .env file if you have environment variables
COPY .env ./

# Expose port 3000 (or whichever port your app runs on)
EXPOSE 3000

# Command to run the Go app
CMD ["./main"]
