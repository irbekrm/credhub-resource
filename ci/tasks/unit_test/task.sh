#!/bin/bash

set -euox pipefail

cd credhub_resource

go test ./...
