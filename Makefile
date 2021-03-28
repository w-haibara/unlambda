unlambda: *.go cmd/unlambda/*.go cli/*.go
	gofmt -w *.go cmd/unlambda/*.go cli/*.go
	go build ./cmd/...

.PHONY: init
init:
	go mod init unlambda
	go mod tidy

.PHONY: test
test:
	gofmt -w *.go cmd/unlambda/*.go cli/*.go
	go test -v ./...
