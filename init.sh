#!/bin/bash

echo "========Installing Go tools..."
GO111MODULE=on go install golang.org/x/tools/gopls@latest
GO111MODULE=on go install github.com/cweill/gotests/...@latest
GO111MODULE=on go install github.com/fatih/gomodifytags@latest
GO111MODULE=on go install github.com/josharian/impl@latest
GO111MODULE=on go install github.com/haya14busa/goplay/cmd/goplay@latest
GO111MODULE=on go install github.com/go-delve/delve/cmd/dlv@latest
GO111MODULE=on go install honnef.co/go/tools/cmd/staticcheck@latest

echo "========Go tools installation completed."

echo "========Initializing Go module..."
go mod init golang-learning

echo "========Go module initialized."

echo "========Downloading dependencies..."
go mod tidy

echo "========Dependencies downloaded."

echo "========Go project initialization completed."
