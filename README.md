# Credhub Resource

A [concourse](https://concourse-ci.org/) resource for [Credhub](https://github.com/cloudfoundry-incubator/credhub). 

Intended to be used for retrieving a credential that cannot be retrieved by Concourse in the [standard way](https://concourse-ci.org/credhub-credential-manager.html)

:warning: The credential's value will be retrieved from CredHub and stored in the resource container in plain text. :warning:


## Source Configuration

* `server`: *Required.* The address of the Credhub server
* `client_name`: *Required.* The UAA client ID for authorizing with Credhub.
* `client_secret`: *Required.* The UAA client secret for authorizing with Credhub.
* `name`: *Required.* The credential name to be retrieved.
* `skip_tls_validation`: *Optional.* Skip TLS validation for connections to Credhub and UAA. (At the moment it is *not* possible to pass CredHub CA certs to this resource)

### Source example

``` yaml

resource_types:
- name: credhub
  type: docker-image
  source:
    repository: engineerbetter/concourse-credhub-resource
    
resources:
- name: some_secret
  type: credhub
  source:
    server: https://credhub.example.com
    client_name : admin
    client_secret: admin
    name: /bosh/cf/some_secret
    skip_tls_validation: true
```

## Behaviour

### `check`: Check for new versions of the credential

Check will return the [CredHub ID](https://credhub-api.cfapps.io/version/2.5/#_credential_ids) of the latest version of the credential.

A new CredHub ID is generated every time the credential is set, even if its value remains the same.

### `in` Retrieve the credential

Retrieves the credential by the ID that `check` fetched and writes it to `value` file in the destination directory.

### `out` Not implemented

#### Files populated

* `value` Credential value in plain text. Credentials of type `value` and `password` will be stored as a simple JSON string. Credentials of type `json` will be stored as a JSON object or array depending on what the value is.

* `created_at` When the credential was created in CredHub

* `type` [CredHub credential type](https://docs.cloudfoundry.org/credhub/credential-types.html)

* `id` [CredHub credential id](https://credhub-api.cfapps.io/version/2.5/#_credential_ids)

### Usage example

``` yaml
jobs:
- name: test
  plan:
  - get: some_secret
    trigger: true
  - task: test
    image: some_image
    config:
      platform: linux
    inputs: some_secret
    run:
      path: bash
      args:
      - -c
      - |
        secret=$(jq -r '.' some_secret/value)
        # do something with the secret
```
