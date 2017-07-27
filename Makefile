SHELL_IMAGE=golang:1.8.3
GIT_SHA=$(shell git rev-parse --verify HEAD | cut -c1-6)
VERSION=$(shell cat VERSION)
LDFLAGS="-X github.com/serainville/inquisitor/variables.Version=$(VERSION) -X github.com/serainville/inquisitor/variables.CommitID=$(GIT_SHA)"

DOCKER_BUILD=$(shell pwd)/.docker_build
DOCKER_CMD=$(DOCKER_BUILD)/inquisitor


default: clean deps compile

all: clean deps lint gofmt gotest build-linux-amd64 build-darwin-amd64 build-freebsd-amd64 build-windows-amd64

compile:
	go build -o bin/inquisitor-${VERSION}-amd64 -ldflags "-X inquisitor/variables.Version=${VERSION}"

$(DOCKER_CMD): clean_docker
	mkdir -p $(DOCKER_BUILD)
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o $(DOCKER_CMD) .

docker: $(DOCKER_CMD)
	docker build -t serainville/inquisitor:latest .
	docker build -t serainville/inquisitor:${VERSION} .

heroku: $(DOCKER_CMD)
	heroku container:push web


deps:
	go get github.com/spf13/cobra
	go get github.com/inconshreveable/mousetrap
	go get github.com/shirou/gopsutil

clean:
	rm -rf bin/*

clean_docker:
	rm -rf $(DOCKER_BUILD)

build:
	go build -o bin/inquisitor-${VERSION}-amd64 -ldflags $(LDFLAGS)

lint:
	$(GOPATH)/bin/golint ./...

gofmt:
	gofmt -s -r .

gotest:
	go test

build-linux-amd64:
	GOOS=linux GOARCH=amd64 go build -v -o bin/linux_amd64/inquisitor-${VERSION}-amd64 -ldflags $(LDFLAGS)

build-darwin-amd64:
	GOOS=darwin GOARCH=amd64 go build -v -o bin/darwin_amd64/inquisitor-${VERSION}-amd64 -ldflags $(LDFLAGS)

build-freebsd-amd64:
	GOOS=freebsd GOARCH=amd64 go build -v -o bin/freebsd_amd64/inquisitor-${VERSION}-amd64 -ldflags $(LDFLAGS)

build-windows-amd64:
	GOOS=windows GOARCH=amd64 go build -v -o bin/windows_amd64/inquisitor-${VERSION}-amd64 -ldflags $(LDFLAGS)
