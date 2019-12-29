#!/bin/bash

set -ex 

# Set fly target and fill example_vars.yml with actual values
target=\<your-fly-target\>

# Log into credhub
# Set some test secrets
credhub set -n /some/random/path/to/secret -t value -v secret_value
credhub set -n /some/user -t json -v '{"username":"Jane","password":"Doe"}'

# Set the test pipeline
fly -t $target set-pipeline \
  --pipeline=test_credhub_resource \
  --config=example_pipeline.yml \
  --var=test_secret_name="/some/random/path/to/secret" \
  --var=test_secret_user_name="/some/user" \
  --load-vars-from=example_vars.yml \
  --check-creds \
  --non-interactive

fly -t $target unpause-pipeline \
  --pipeline=test_credhub_resource
