# Use an official Golang image as the base
FROM golang:1.21

# Set environment variables
ENV GO111MODULE=on \
    CGO_ENABLED=1 \
    GOOS=linux \
    GOARCH=amd64

# Install wkhtmltopdf dependencies
RUN apt-get update && \
    apt-get install -y wkhtmltopdf libxrender1 libfontconfig1 libx11-dev && \
    rm -rf /var/lib/apt/lists/*

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire project
COPY . .

# Build the Go app
RUN go build -o invoicer main.go

# Expose the app port
EXPOSE 8080

# Command to run the executable
CMD ["./invoicer"]
