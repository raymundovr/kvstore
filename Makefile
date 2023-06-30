.PHONY: test

BINARY_NAME=kvs

all: test build

build:
	go build -a -o ./bin/${BINARY_NAME}

clean:
	go clean
	rm ./bin/${BINARY_NAME}

test:
	go test -v ./...