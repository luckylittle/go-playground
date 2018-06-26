# GO LANG

## Introduction

1. Configure `$GOPATH`, e.g. `mkdir -p ~/Projects/go/{bin,pkg,src};mkdir -p ~/Projects/go/src/github.com/luckylittle;echo export 'GOPATH=/home/lmaly/Projects/go' >> ~/.zshrc`
2. Write Hello World in `main.go` in `$GOPATH/src/github.com/luckylittle/hello/main.go`
3. `go build` and `./hello` inside `$GOPATH/src/github.com/luckylittle/hello/`, cleanup with `go clean`
    or `go run main.go` which does not build `./hello` so you don't need to do `go clean`

## Types
