# Start from golang base image
FROM golang:alpine as builder

# ENV GO111MODULE=on

# Set the current working directory inside the container
WORKDIR /app

RUN apk add --no-cache bash

# Copy the source from the current directory to the working Directory inside the container
COPY . .

# Build the Go app
# RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

RUN bash setup.sh
RUN go get github.com/pressly/goose/cmd/goose
RUN bash migrate.sh up

# Start a new stage from scratch
FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage. Observe we also copied the .env file
COPY --from=builder /app/main .
COPY --from=builder /app/.env .

# Expose port 8080 to the outside world
EXPOSE 8080

#Command to run the executable
CMD ["./main"]