# Use golang image as the base
FROM golang:1.22.5

# Additional image metadata
LABEL version="1.0"
LABEL maintainer="johneliud4@gmail.com"
LABEL description="An image file for a web server that accepts text as input and displays the output using ASCII characters."

# Specify the working directory of the image
WORKDIR /app

# Copy go.mod file
COPY go.mod ./

# Download dependencies
RUN go mod download

# Copy the rest of the source code
COPY . .

# Command to build an executable
RUN go build -o ascii-art-web .

# Expose image to the outside world using port 8080
EXPOSE 8080

# Run the specified executable file
CMD ["./ascii-art-web"]