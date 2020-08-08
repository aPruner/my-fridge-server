# Start from golang base image
FROM golang:alpine as builder

# ENV GO111MODULE=on

# Set the current working directory inside the container
WORKDIR /app

# Copy the source from the current directory to the working Directory inside the container
COPY . .

# Get dependencies
RUN go get

# Build the go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]