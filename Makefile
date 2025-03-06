.PHONY: build-example run-example build-node run-node

build-example:
	go build -o bin/example/main cmd/example/main.go

run-example:
	go run cmd/example/main.go


build-node:
	go build -o bin/node/main cmd/node/main.go

run-node:
	go run cmd/node/main.go

