package credhub

import (
	"code.cloudfoundry.org/credhub-cli/credhub/credentials"
)

//go:generate mockgen -source $GOFILE -destination ../mocks/credhub.go -package mocks
type CredHub interface {
	GetLatestVersion(string) (credentials.Credential, error)
	GetById(string) (credentials.Credential, error)
}
