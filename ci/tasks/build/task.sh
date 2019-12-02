#!/bin/bash

set -euox pipefail

export BUILT_BINARIES_DIR=$PWD/built-binaries

cd credhub_resource

go build -mod=vendor -o "$BUILT_BINARIES_DIR/out" ./cmd/out
go build -mod=vendor -o "$BUILT_BINARIES_DIR/in" ./cmd/in
go build -mod=vendor -o "$BUILT_BINARIES_DIR/check" ./cmd/check

cp Dockerfile "$BUILT_BINARIES_DIR"
