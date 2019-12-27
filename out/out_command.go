package out

import (
	"github.com/EngineerBetter/credhub-resource/concourse"
	"github.com/EngineerBetter/credhub-resource/credhub"
)

type OutResponse struct {
	Version  concourse.Version    `json:"version"`
	Metadata []concourse.Metadata `json:"metadata"`
}

type OutCommand struct {
	client             credhub.CredHub
	resourcesDirectory string
}

func NewOutCommand(client credhub.CredHub, resourcesDirectory string) OutCommand {
	return OutCommand{
		client:             client,
		resourcesDirectory: resourcesDirectory,
	}
}

func (c OutCommand) Run(outRequest concourse.OutRequest) (OutResponse, error) {
	concourseOutput := OutResponse{
		Version:  concourse.Version{},
		Metadata: []concourse.Metadata{},
	}
	latest, err := c.client.GetLatestVersion(outRequest.Source.Path)
	if err != nil {
		return concourseOutput, err
	}

	version := concourse.Version{
		Server: outRequest.Source.Server,
		ID:     latest.Metadata.Id,
	}
	concourseOutput.Version = version

	return concourseOutput, nil
}
