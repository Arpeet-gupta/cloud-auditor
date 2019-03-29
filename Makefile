all: get-depends build-code

get-depends:
	go get -v -t -d ./...

build-code:
	go build

