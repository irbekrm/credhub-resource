package check

import (
	"github.com/EngineerBetter/credhub-resource/concourse"
	"github.com/EngineerBetter/credhub-resource/credhub"
)

type CheckCommand struct {
	client credhub.CredHub
}

func NewCheckCommand(client credhub.CredHub) CheckCommand {
	return CheckCommand{client: client}
}

func (c CheckCommand) Run(checkRequest concourse.CheckRequest) ([]concourse.Version, error) {
	latest, err := c.client.GetLatestVersion(checkRequest.Source.Name)

	if err != nil {
		return []concourse.Version{concourse.Version{}}, err
	}

	version := concourse.Version{
		ID:     latest.Metadata.Id,
		Server: checkRequest.Source.Server,
	}

	return []concourse.Version{version}, nil
}
