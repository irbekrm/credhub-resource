package concourse

import (
	"bytes"
	"encoding/json"
	"fmt"
)

const MissingServer = "MISSING-SERVER-SHORTCIRCUIT.example.com"

type InRequest struct {
	Source  Source   `json:"source"`
	Version Version  `json:"version"`
	Params  InParams `json:"params"`
}

func NewInRequest(request []byte) (InRequest, error) {
	var inRequest InRequest
	if err := json.NewDecoder(bytes.NewReader(request)).Decode(&inRequest); err != nil {
		return InRequest{}, fmt.Errorf("invalid parameters: %s", err)
	}

	if inRequest.Source.Server == "" {
		inRequest.Source.Server = MissingServer
	}

	return inRequest, nil
}
