# Start from golang base image
FROM golang:alpine

# Get build-base and git so that goose will work
RUN apk update && apk add build-base
RUN apk update && apk add --no-cache git

# Get goose
RUN go get -u github.com/pressly/goose/cmd/goose

# Run the migrations
CMD ["sh", "-c", "goose postgres \"host=${DB_HOSTNAME}:5432 user=postgres password=adminpass1234 dbname=myfridge sslmode=disable\" up"]