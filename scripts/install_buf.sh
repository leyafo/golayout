#!/bin/sh

# Substitute GOBIN for your bin directory
# Leave unset to default to $GOPATH/bin
GO111MODULE=on GOBIN=$PWD/tools go get \
    github.com/bufbuild/buf/cmd/buf \
    github.com/bufbuild/buf/cmd/protoc-gen-buf-breaking \
    github.com/bufbuild/buf/cmd/protoc-gen-buf-lint \

GO111MODULE=on go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
