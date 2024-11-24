
BINARY_NAME := ip2country
DOCKER_IMAGE := ip2country
TAG ?= latest

build:
	go build -o $(BINARY_NAME)

test:
	go test ./...

run:
	go run .

clean:
	rm ./$(BINARY_NAME)

build-docker:
	docker build -t $(DOCKER_IMAGE):$(TAG) .
