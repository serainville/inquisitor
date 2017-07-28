SHELL_IMAGE=golang:1.8.3
GIT_SHA=$(shell git rev-parse --verify HEAD | cut -c1-6)
VERSION=$(shell cat VERSION)
LDFLAGS="-X github.com/serainville/inquisitor/variables.Version=$(VERSION) -X github.com/serainville/inquisitor/variables.CommitID=$(GIT_SHA)"

DOCKER_BUILD=$(shell pwd)/.docker_build
DOCKER_CMD=$(DOCKER_BUILD)/inquisitor


default: clean deps compile

all: clean deps lint gofmt gotest build-darwin build-freebsd build-windows build-netbsd build-linux

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
	gofmt -s -w .

gotest:
	go test

build-darwin:
	GOOS=darwin GOARCH=amd64 go build -o bin/darwin/inquisitor-${VERSION}-amd64 -ldflags $(LDFLAGS)

build-freebsd:
	GOOS=freebsd GOARCH=amd64 go build -o bin/freebsd/inquisitor-${VERSION}-amd64 -ldflags $(LDFLAGS)

build-windows:
	GOOS=windows GOARCH=amd64 go build -o bin/windows/inquisitor-${VERSION}-amd64.exe -ldflags $(LDFLAGS)
	GOOS=windows GOARCH=386 go build -o bin/windows/inquisitor-${VERSION}-386.exe -ldflags $(LDFLAGS)

build-netbsd:
	GOOS=netbsd GOARCH=386 go build -o bin/netbsd/inquisitor-${VERSION}-386 -ldflags $(LDFLAGS)
	GOOS=netbsd GOARCH=amd64 go build -o bin/netbsd/inquisitor-${VERSION}-amd64 -ldflags $(LDFLAGS)
	GOOS=netbsd GOARCH=arm go build -o bin/netbsd/inquisitor-${VERSION}-arm -ldflags $(LDFLAGS)

build-linux:
	GOOS=linux GOARCH=386 go build -o bin/linux/inquisitor-${VERSION}-386 -ldflags $(LDFLAGS)
	GOOS=linux GOARCH=amd64 go build -o bin/linux/inquisitor-${VERSION}-amd64 -ldflags $(LDFLAGS)
	GOOS=linux GOARCH=arm go build -o bin/linux/inquisitor-${VERSION}-arm -ldflags $(LDFLAGS)
	GOOS=linux GOARCH=arm64 go build -o bin/linux/inquisitor-${VERSION}-arm64 -ldflags $(LDFLAGS)






