#!/bin/bash

set -eux

yes | apt-get install shellcheck

# Installs golangci-lint https://github.com/golangci/golangci-lint
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.21.0

cd credhub-resource

# Runs recursively
golangci-lint run --skip-dirs-use-default -E stylecheck -E bodyclose -E dupl -E gochecknoglobals -E gochecknoinits -E goconst

# Globally ignoring SC2154 as it doesn't play nice with variables
# set by Concourse for use in tasks.
# https://github.com/koalaman/shellcheck/wiki/SC2154
find . -name vendor -prune ! -type d -o -name '*.sh' -exec shellcheck -e SC2154 {} +