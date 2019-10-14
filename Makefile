export GO111MODULE=on

.PHONY: deps
deps:
	go mod tidy

.PHONY: build
build: deps
	go build -o gtc

.PHONY: test
test: deps
	go test ./...

.PHONY: lint
lint: deps
	go vet ./...
	golint -set_exit_status ./...
