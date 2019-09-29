all: build test

build:
	go build raytracer/main.go

test:
	go test ./...
