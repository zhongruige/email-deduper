FROM golang:1.14-alpine as build

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

RUN apk add build-base
RUN apk add --no-cache git

RUN mkdir /email-deduper
ADD . /email-deduper
WORKDIR /email-deduper

RUN go mod download
RUN go build -o main

CMD ["/email-deduper/main"]
