all: get-depends build-code

get-depends:
	go get -t -v ./...
	go install ./...

build-code:
	go build

