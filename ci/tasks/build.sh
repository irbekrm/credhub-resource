#!/usr/bin/env bash

set -eux

export BUILT_BINARIES_DIR=$PWD/built-binaries

cd credhub-resource

go build -mod=vendor -o "$BUILT_BINARIES_DIR/out" ./cmd/out
go build -mod=vendor -o "$BUILT_BINARIES_DIR/in" ./cmd/in
go build -mod=vendor -o "$BUILT_BINARIES_DIR/check" ./cmd/check

cp Dockerfile "$BUILT_BINARIES_DIR"
