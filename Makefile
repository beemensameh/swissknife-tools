all:build

build:build-linux build-win32 build-win64 build-mac

build-linux:
	go build -mod=vendor -o=./bin/swisstool -race ./cmd/swissknife-tools/.

build-win32:
	GOOS=windows GOARCH=386 go build -mod=vendor -o=./bin/swisstool.exe ./cmd/swissknife-tools/.

build-win64:
	GOOS=windows GOARCH=amd64 go build -mod=vendor -o=./bin/swisstool64.exe ./cmd/swissknife-tools/.

build-mac:
	GOOS=darwin GOARCH=amd64 go build -mod=vendor -o=./bin/swisstool-mac ./cmd/swissknife-tools/.

tidy:
	go mod vendor
