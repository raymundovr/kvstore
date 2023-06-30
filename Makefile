.PHONY: test docker

BINARY_NAME=kvs

all: test build

build:
	go build -a -o ./bin/${BINARY_NAME}

clean:
	go clean
	rm ./bin/${BINARY_NAME}

docker:
	docker build --tag kvs:latest .

test:
	go test -v ./...