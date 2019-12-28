package concourse

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

type Source struct {
	Server            string `json:"server,omitempty" yaml:"server"`
	ClientName        string `json:"client_name,omitempty" yaml:"client_name"`
	ClientSecret      string `json:"client_secret,omitempty" yaml:"client_secret"`
	Name              string `json:"name,omitempty" yaml:"name"`
	SkipTLSValidation bool   `json:"skip_tls_validation,omitempty" yaml:"skip_tls_validation"`
}

type sourceRequest struct {
	Source Source `json:"source"`
}

func NewSource(config []byte, sourcesDir string) (Source, error) {
	var sourceRequest sourceRequest
	if err := json.NewDecoder(bytes.NewReader(config)).Decode(&sourceRequest); err != nil {
		return Source{}, fmt.Errorf("invalid dynamic source config: %s", err)
	}

	if err := checkRequiredSourceParameters(sourceRequest.Source); err != nil {
		return Source{}, err
	}

	return sourceRequest.Source, nil
}

func checkRequiredSourceParameters(source Source) error {
	missingParameters := []string{}

	if source.Server == "" {
		missingParameters = append(missingParameters, "server")
	}
	if source.ClientName == "" {
		missingParameters = append(missingParameters, "client_name")
	}
	if source.ClientSecret == "" {
		missingParameters = append(missingParameters, "client_secret")
	}

	if len(missingParameters) > 0 {
		parametersString := "parameter"
		if len(missingParameters) > 2 {
			parametersString = parametersString + "s"
		}
		errorMessage := fmt.Sprintf("Missing required source %s: %s", parametersString, strings.Join(missingParameters, ", "))
		return errors.New(errorMessage)
	}

	return nil
}
