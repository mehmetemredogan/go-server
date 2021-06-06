-include .env

VERSION=$(shell git describe --tags)
BUILD=$(shell git rev-parse --short HEAD)
PROJECT_NAME=go-server

LDFLAGS=-ldflags "-X=main.Version=$(VERSION) -X=main.Build=$(BUILD)"

build:
	@echo "  |>  Building..."
	go build $(LDFLAGS) -o bin/$(PROJECT_NAME).out cmd/$(PROJECT_NAME)/main.go

clean:
	@echo "  |>  Cleaning..."
	go clean

get:
	@echo "  |>  Checking missing dependencies... "
	go get $(get)

run:
	@echo "  |>  Running..."
	./bin/$(PROJECT_NAME).out

tidy:
	@echo "  |>  Dependency removing..."
	go mod tidy -v

dev: clean build run

.PHONY: build test clean