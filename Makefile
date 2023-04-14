all:build

build:build-linux32 build-linux64 build-win32 build-win64 build-mac

test:
	go test ./... -timeout 30s

test-verbose:
	go test ./... -race -covermode=atomic -timeout 30s -v

test-coverage:
	go test ./... -race -covermode=atomic -timeout 30s -coverprofile=coverage.out

build-linux32:
	GOOS=linux GOARCH=386 go build -mod=vendor -o=./bin/swisstool ./cmd/swissknife-tools/.

build-linux64:
	GOOS=linux GOARCH=amd64 go build -mod=vendor -o=./bin/swisstool-amd64 -race ./cmd/swissknife-tools/.

build-win32:
	GOOS=windows GOARCH=386 go build -mod=vendor -o=./bin/swisstool-i386.exe ./cmd/swissknife-tools/.

build-win64:
	GOOS=windows GOARCH=amd64 go build -mod=vendor -o=./bin/swisstool-x64.exe ./cmd/swissknife-tools/.

build-mac:
	GOOS=darwin GOARCH=amd64 go build -mod=vendor -o=./bin/swisstool-darwin-amd64 ./cmd/swissknife-tools/.

vendor:
	go mod tidy && go mod vendor

vulncheck:
	govulncheck ./...

lint:
	golangci-lint run
