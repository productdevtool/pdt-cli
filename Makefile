GOOS ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)

.PHONY: all build clean

all: build

build:
	go build -o pdt-$(GOOS)-$(GOARCH) .

clean:
	rm -f pdt-*-*
