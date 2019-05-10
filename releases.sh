#!/bin/bash
env GOOS=darwin GOARCH=amd64 go build -o build/barge-darwin-amd64 .
env GOOS=linux GOARCH=386 go build -o build/barge-linux-386 .
env GOOS=linux GOARCH=amd64 go build -o build/barge-linux-amd64 .
env GOOS=windows GOARCH=amd64 go build -o build/barge-windows-amd64.exe .