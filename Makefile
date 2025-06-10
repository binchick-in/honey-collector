VERSION ?= 1.2

build: build-mac build-linux build-windows

build-mac:
	mkdir -p build
	GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o build/honey-collector-mac-amd64-$(VERSION) ./cmd/honey-collector
	GOOS=darwin GOARCH=arm64 go build -ldflags="-s -w" -o build/honey-collector-mac-arm64-$(VERSION) ./cmd/honey-collector

build-linux:
	mkdir -p build
	GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o build/honey-collector-linux-amd64-$(VERSION) ./cmd/honey-collector
	GOOS=linux GOARCH=arm64 go build -ldflags="-s -w" -o build/honey-collector-linux-arm64-$(VERSION) ./cmd/honey-collector

build-windows:
	mkdir -p build
	GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o build/honey-collector-windows-amd64-$(VERSION).exe ./cmd/honey-collector
	GOOS=windows GOARCH=arm64 go build -ldflags="-s -w" -o build/honey-collector-windows-arm64-$(VERSION).exe ./cmd/honey-collector

format:
	@go fmt ./...
