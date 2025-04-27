# Use the official Golang image as the base image
FROM golang:1.23 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the Go application with static linking
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o myapp ./cmd/web

# Use a minimal base image for the production stage
FROM alpine:latest

# Install necessary libraries if required
RUN apk add --no-cache ca-certificates

# Install OpenSSL
RUN apk add --no-cache openssl

# Display the generated files (for demonstration purposes)
RUN ls -l

# Set the working directory inside the container
WORKDIR /root/

COPY ./ui ./ui
COPY app_config.json .

# Generate the SSL key pair
RUN openssl genpkey -algorithm RSA -out private.key && \
    openssl req -new -key private.key -out csr.csr -subj "/CN=example.com" && \
    openssl x509 -req -days 365 -in csr.csr -signkey private.key -out dev_cert.crt

# Copy the built Go application from the builder stage
COPY --from=builder /app/myapp .

# Define a build argument for the build date and time
ARG BUILD_DATE

# Set the build date and time as an environment variable
ENV BUILD_DATE=$BUILD_DATE

RUN echo "Build date and time: $BUILD_DATE"

# Expose the port the application runs on
EXPOSE 8442

# Command to run the application
CMD ["./myapp"]