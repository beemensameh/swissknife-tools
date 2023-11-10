VERSION := $(shell git rev-parse --abbrev-ref HEAD)
OS = $(shell go env GOOS)
ARCH = $(shell go env GOARCH)

ifneq ($(OS), windows)
	ifeq ($(ARCH), amd64)
		RACE = -race
	endif
endif

all:build-linux build-win build-mac
build-linux:build-linux32 build-linux64
build-win:build-win32 build-win64
build-mac:build-mac-amd build-mac-arm

build:
	GOOS=$(OS) GOARCH=$(ARCH) go build -ldflags "-s -X main.name=swisstools-$(ARCH) -X main.version=$(VERSION)" -mod=vendor -o=./bin/swisstools-$(ARCH) $(RACE) ./cmd/swissknife-tools/.

test:
	go test ./... -timeout 30s

test-verbose:
	go test ./... -race -covermode=atomic -timeout 30s -v

test-coverage:
	go test ./... -race -covermode=atomic -timeout 30s -coverprofile=coverage.out

build-linux32:
	GOOS=linux GOARCH=386 go build -ldflags "-s -X main.name=swisstools -X main.version=$(VERSION)" -mod=vendor -o=./bin/swisstools ./cmd/swissknife-tools/.

build-linux64:
	GOOS=linux GOARCH=amd64 go build -ldflags "-s -X main.name=swisstools-amd64 -X main.version=$(VERSION)" -mod=vendor -o=./bin/swisstools-amd64 -race ./cmd/swissknife-tools/.

build-win32:
	GOOS=windows GOARCH=386 go build -ldflags "-s -X main.name=swisstools-i386.exe -X main.version=$(VERSION)" -mod=vendor -o=./bin/swisstools-i386.exe ./cmd/swissknife-tools/.

build-win64:
	GOOS=windows GOARCH=amd64 go build -ldflags "-s -X main.name=swisstools-x64.exe -X main.version=$(VERSION)" -mod=vendor -o=./bin/swisstools-x64.exe ./cmd/swissknife-tools/.

build-mac-amd:
	GOOS=darwin GOARCH=amd64 go build -ldflags "-s -X main.name=swisstools-darwin-amd64 -X main.version=$(VERSION)" -mod=vendor -o=./bin/swisstools-darwin-amd64 -race ./cmd/swissknife-tools/.

build-mac-arm:
	GOOS=darwin GOARCH=arm64 go build -ldflags "-s -X main.name=swisstools-darwin-arm64 -X main.version=$(VERSION)" -mod=vendor -o=./bin/swisstools-darwin-arm64 ./cmd/swissknife-tools/.

vendor:
	go mod tidy && go mod vendor

vulncheck-install:
	go install golang.org/x/vuln/cmd/govulncheck@latest

vulncheck:
	govulncheck ./...

lint:
	golangci-lint run

auto-complete:
	source <(./bin/swisstool completion bash)
