# Start from golang base image
FROM golang:alpine

# Get build-base and git so that goose will work
RUN apk update && apk add build-base
RUN apk update && apk add --no-cache git

RUN go get -u github.com/pressly/goose/cmd/goose

# May need to also pass postgres conn string to goose here

CMD ["goose", "postgres=\"user=adam password=adminpass1234 dbname=myfridge sslmode=disable\"", "up"]