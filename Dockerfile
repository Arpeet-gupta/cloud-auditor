FROM golang:latest

MAINTAINER Abhishek Dubey

COPY ./ /go/src/github.com/iamabhishek-dubey/

RUN go get -v -t -d ./... \
    && go build

