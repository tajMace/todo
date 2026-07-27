BINARY_NAME := todo
INSTALL_PATH := $(HOME)/go/bin/$(BINARY_NAME)

.PHONY: install run test

install:
	go build -o $(INSTALL_PATH) ./src

run:
	go run ./src $(ARGS)

test:
	go test ./...
