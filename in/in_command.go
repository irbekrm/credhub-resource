package in

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/EngineerBetter/credhub-resource/concourse"
	"github.com/EngineerBetter/credhub-resource/credhub"
)

type InCommand struct {
	client credhub.CredHub
}

type InResponse struct {
	Version  concourse.Version    `json:"version"`
	Metadata []concourse.Metadata `json:"metadata"`
}

func NewInCommand(client credhub.CredHub) InCommand {
	return InCommand{client: client}
}

func (c InCommand) Run(inRequest concourse.InRequest, targetDir string) (InResponse, error) {
	credential, err := c.client.GetById(inRequest.Version.ID)
	if err != nil {
		return InResponse{}, err
	}
	valueToBytes, err := json.Marshal(credential.Value)
	if err != nil {
		return InResponse{}, err
	}
	err = createFile("value", targetDir, valueToBytes)
	if err != nil {
		return InResponse{}, err
	}

	err = createFile("created_at", targetDir, []byte(credential.Metadata.Base.VersionCreatedAt))
	if err != nil {
		return InResponse{}, err
	}

	err = createFile("type", targetDir, []byte(credential.Metadata.Type))
	if err != nil {
		return InResponse{}, err
	}

	err = createFile("id", targetDir, []byte(credential.Metadata.Id))
	if err != nil {
		return InResponse{}, err
	}

	actualVersion := concourse.Version{
		ID:     credential.Metadata.Id,
		Server: inRequest.Source.Server,
	}

	metadata := concourse.Metadata{
		ID:               credential.Metadata.Id,
		VersionCreatedAt: credential.Metadata.Base.VersionCreatedAt,
		Type:             credential.Metadata.Type,
	}

	if actualVersion.Server != inRequest.Version.Server {
		return InResponse{}, errors.New("credhub server is different than configured source")
	}

	return InResponse{Version: actualVersion, Metadata: []concourse.Metadata{metadata}}, nil
}

func createFile(filename, targetDir string, data []byte) error {
	filepath := fmt.Sprintf("%s/%s", targetDir, filename)
	err := ioutil.WriteFile(filepath, data, 0644)
	return err
}
