# Base image with Go
FROM golang:tip-trixie

# Create app directory
WORKDIR /app

# Copy project files
COPY . .

# Build the Go program
RUN go build -o server server.go

EXPOSE 6700

# Run the binary
CMD ["/app/server"]




