# Use the official Golang image as the base image
FROM golang:1.19 as builder

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum files to the container
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the source code to the container
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o beprayed-app-worker .

# Use the official Alpine Linux as the base image for the final image
FROM alpine:latest

# Add ca-certificates to the final image
RUN apk --no-cache add ca-certificates

# Set the working directory
WORKDIR /root/

# Copy the built binary from the builder stage
COPY --from=builder /app/beprayed-app-worker .

ARG ENV_FILE=.env.prod
# Copy the .env file to the container
COPY ${ENV_FILE} ./.env

# Expose the application port
EXPOSE 8080

# Run the application
CMD ["./beprayed-app-worker"]
