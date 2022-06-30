#!/bin/bash
GOOS=windows GOARCH=amd64 go build -o bin/cc-rpc-win-amd64.exe main.go
GOOS=linux GOARCH=amd64 go build -o bin/cc-rpc-linux-amd64 main.go
GOOS=darwin GOARCH=amd64 go build -o bin/cc-rpc-macos-amd64 main.go
GOOS=darwin GOARCH=arm64 go build -o bin/cc-rpc-macos-arm64 main.go