.DEFAULT_GOAL := build

BUILD_REVISOIN := $(shell git describe)
BUILD_DATE := $(shell date -u +"%Y-%m-%d")
BUILD_TIME := $(shell date -u +"%H:%M:%S")
BUILD_USER := $(shell whoami)
BUILD_MACHINE := $(shell hostname)
release: export CGO_ENABLED = 0
release: export GO111MODULE = on

PACKAGE := github.com/vitanovs/versioner
VERSION_PACKAGE := ${PACKAGE}/version

BINARY_NAME := versioner

LDFLAGS = -ldflags \
           "-s -w \
            -X ${VERSION_PACKAGE}.BuildRevision=${BUILD_REVISOIN} \
            -X ${VERSION_PACKAGE}.BuildDate=${BUILD_DATE} \
            -X ${VERSION_PACKAGE}.BuildTime=${BUILD_TIME} \
            -X ${VERSION_PACKAGE}.BuildUser=${BUILD_USER} \
            -X ${VERSION_PACKAGE}.BuildMachine=${BUILD_MACHINE}"

build: mod
	go build -mod vendor ${LDFLAGS} -o ./bin/${BINARY_NAME}

release: build

docker:
	docker build -t ${BINARY_NAME} .
	docker image prune --filter label=stage=versioner-builder-env -f

clean:
	rm -f ./bin/${BINARY_NAME}

install: build
	install -m 0755 ./bin/${BINARY_NAME} ${GOPATH}/bin/${BINARY_NAME}

uninstall:
	rm -rf ${GOPATH}/bin/${BINARY_NAME}

darwin:
	GOOS=darwin GOARCH=amd64 go build -mod vendor ${LDFLAGS} -o ./bin/${BINARY_NAME}
	tar cvfz ${BINARY_NAME}-darwin-amd64-latest.tar.gz ./bin/${BINARY_NAME}

linux:
	GOOS=linux GOARCH=amd64 go build -mod vendor ${LDFLAGS} -o ./bin/${BINARY_NAME}
	tar cvfz ${BINARY_NAME}-linux-amd64-latest.tar.gz ./bin/${BINARY_NAME}

windows:
	GOOS=windows GOARCH=amd64 go build -mod vendor ${LDFLAGS} -o ./bin/${BINARY_NAME}
	tar cvfz ${BINARY_NAME}-windows-amd64-latest.tar.gz ./bin/${BINARY_NAME}

pack: darwin linux windows

format:
	go fmt ./

mod:
	go mod tidy -v
	go mod vendor -v
	go mod verify

.PHONY: build install uninstall clean format
