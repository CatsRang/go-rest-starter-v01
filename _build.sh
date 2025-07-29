#!/bin/bash

# go mod tidy
# Get current date in YYYYMMDD format
BUILD_DATE=$(date +%Y%m%d)

GOOS=darwin GOARCH=arm64 go build -ldflags "-s -w" -o bin/go-rest_01-darwin_arm64-${BUILD_DATE} cmd/main.go

GOOS=linux GOARCH=amd64 go build  -ldflags "-s -w" -o bin/go-rest_01-linux_amd64-${BUILD_DATE} cmd/main.go

GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o bin/go-rest_01-windows_amd64-${BUILD_DATE}.exe cmd/main.go
