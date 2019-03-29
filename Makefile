all: get-depends build-code build-image

get-depends:
	go get -v -t -d ./...

build-code:
	go build

build-image:
	docker build -t cloud-auditor:latest -f Dockerfile .

