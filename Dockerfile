# Use the official Golang image to create a build artifact.
FROM golang:1.20 as builder

# Copy local code to the container image.
WORKDIR /app
COPY . .

# Build the binary.
RUN go build -o server

# Use a minimal alpine image for the final container
FROM alpine:3.14
WORKDIR /app

# Copy the binary from the builder stage.
COPY --from=builder /app/server /app/server

# Run the server.
CMD ["/app/server"]
