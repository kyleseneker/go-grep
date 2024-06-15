.PHONY: build fmt clean lint

OS := $(shell go env GOOS)
BUILDCMD=env GOOS=$(OS) GOARCH=amd64 go build -v

build:
	@$(BUILDCMD) -o go-grep cmd/go-grep/*.go 

fmt:
	@go fmt ./...

clean:
	@go clean ./...
	@rm -rf ./go-grep

lint:
	@golangci-lint run ./...

benchmark:
	@hyperfine --warmup 2 './go-grep -r "Nirvana" examples/*' 'grep -r "Nirvana" examples/*'
	@$(MAKE) clean
