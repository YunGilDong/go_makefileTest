  
.PHONY: fmt build test bench
.EXPORT_ALL_VARIABLES:

GO111MODULE ?= off

all: build test

test:
	go test -v

build:
	go build -o ../bin/sayhi sayhi.go

run:
	go run sayhi.go