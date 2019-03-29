FROM golang:latest as builder
MAINTAINER Abhishek Dubey
COPY ./ /go/src/github.com/iamabhishek-dubey/cloud-auditor/
WORKDIR /go/src/github.com/iamabhishek-dubey/cloud-auditor/
RUN go get -v -t -d ./... \
    && go build

FROM scratch
MAINTAINER Abhishek Dubey
COPY --from=builder /go/src/github.com/iamabhishek-dubey/cloud-auditor/cloud-auditor /bin/
COPY --from=builder /go/src/github.com/iamabhishek-dubey/cloud-auditor/htmlreports /cloud-auditor/
WORKDIR /cloud-auditor/

