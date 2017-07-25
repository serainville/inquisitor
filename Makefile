SHELL_IMAGE=golang:1.8.3
GIT_SHA=$(shell git rev-parse --verify HEAD)
VERSION=$(shell cat VERSION)

default: clean deps compile

all: clean deps build-linux-amd64 build-darwin-amd64 build-freebsd-amd64 build-windows-amd64

compile:
	go build -o bin/inquisitor-${VERSION}-amd64 -ldflags "-X inquisitor/constants.Version=${VERSION}"

deps:
	go get github.com/spf13/cobra
	go get github.com/inconshreveable/mousetrap

clean:
	rm -rf bin/*

build:
	go build -o bin/inquisitor-${VERSION}-amd64 -ldflags "-X inquisitor/constants.Version=${VERSION}" 

build-linux-amd64:
	GOOS=linux GOARCH=amd64 go build -v -o bin/linux_amd64/inquisitor-${VERSION}-amd64 -ldflags "-X inquisitor/constants.Version=${VERSION}" 

build-darwin-amd64:
	GOOS=darwin GOARCH=amd64 go build -v -o bin/darwin_amd64/inquisitor-${VERSION}-amd64  -ldflags "-X inquisitor/constants.Version=${VERSION}" 

build-freebsd-amd64:
	GOOS=freebsd GOARCH=amd64 go build -v -o bin/freebsd_amd64/inquisitor-${VERSION}-amd64  -ldflags "-X inquisitor/constants.Version=${VERSION}"

build-windows-amd64:
	GOOS=windows GOARCH=amd64 go build -v -o bin/windows_amd64/inquisitor-${VERSION}-amd64  -ldflags "-X inquisitor/constants.Version=${VERSION}"
