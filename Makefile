VERSION := $(shell git rev-parse --abbrev-ref HEAD)

all:build

build:build-linux32 build-linux64 build-win32 build-win64 build-mac

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

build-mac:
	GOOS=darwin GOARCH=amd64 go build -ldflags "-s -X main.name=swisstools-darwin-amd64 -X main.version=$(VERSION)" -mod=vendor -o=./bin/swisstools-darwin-amd64 ./cmd/swissknife-tools/.

tidy:
	go mod tidy && go mod vendor

vulncheck:
	govulncheck ./...

lint:
	golangci-lint run

auto-complete:
	source <(./bin/swisstool completion bash)
