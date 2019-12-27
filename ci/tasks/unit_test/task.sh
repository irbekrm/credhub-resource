#!/bin/bash

set -euox pipefail

cd credhub_resource

go get -u github.com/golang/mock/mockgen

go generate credhub/credhub.go

go test ./...
