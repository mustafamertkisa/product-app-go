FROM golang:1.22

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go build -o main ./cmd/

# Expose port 8000 to the outside world
EXPOSE 8000

# Command to run the executable
CMD ["./main"]