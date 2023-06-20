# Use the official Go image as the base image
FROM golang:1.20.5-alpine

# Socat needs to be installed for communication with the sidecar
RUN apk add socat
RUN apk add python3

# This directory needs to exist
RUN mkdir -p /codequest

# Set the working directory inside the container
WORKDIR /codequest

# Copy the Go mod and sum files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code into the container
COPY . .
RUN chmod +x /codequest/run.sh

WORKDIR /codequest/cmd/submission

# Build the Go application
RUN go build -o codequest-submission .

# Set the entry point for the container
CMD ["/bin/sh", "-c", "./run.sh"]
