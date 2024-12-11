# Use an official Golang runtime as a parent image
FROM golang:1.23-alpine as builder

# Accept build arguments
ARG SSH_HOST
ARG SSH_USER
ARG SSH_PORT
ARG SSH_PASSWORD
ARG DB_HOST
ARG DB_PORT
ARG LOCAL_PORT

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the Go Modules manifests to the container
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app with build arguments as part of the build process
RUN go build -o main .

# Start a new stage from the smaller Alpine image
FROM alpine:latest

# Install necessary dependencies for SSH (OpenSSH client)
RUN apk --no-cache add openssh-client

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy the compiled binary from the previous stage
COPY --from=builder /app/main .

# Set a default command using build-time arguments
CMD echo "Starting application with the following settings:" \
    && echo "SSH Host: ${SSH_HOST}" \
    && echo "SSH User: ${SSH_USER}" \
    && echo "SSH Port: ${SSH_PORT}" \
    && echo "SSH Password: ${SSH_PASSWORD}" \
    && echo "DB Host: ${DB_HOST}" \
    && echo "DB Port: ${DB_PORT}" \
    && echo "Local Port: ${LOCAL_PORT}" \
    && ./main gw \
    --ssh-host=${SSH_HOST} \
    --ssh-user=${SSH_USER} \
    --ssh-password=${SSH_PASSWORD} \
    --ssh-port=${SSH_PORT} \
    --ssh-auth-type=${SSH_AUTH_TYPE} \
    --ssh-key=${SSH_KEY_PATH} \
    --db-host=${DB_HOST} \
    --db-port=${DB_PORT} \
    --local-host=${LOCAL_HOST} \
    --local-port=${LOCAL_PORT}