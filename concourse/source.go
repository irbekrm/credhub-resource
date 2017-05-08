package concourse

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v2"
)

type Source struct {
	Server            string `json:"server,omitempty" yaml:"server"`
	Username          string `json:"username,omitempty" yaml:"username"`
	Password          string `json:"client_secret,omitempty" yaml:"client_secret"`
	SkipTLSValidation bool   `json:"skip_tls_validation,omitempty" yaml:"skip_tls_validation"`
}

type sourceRequest struct {
	Source Source              `json:"source"`
	Params dynamicSourceParams `json:"params"`
}

type dynamicSourceParams struct {
	SourceFile string `json:"source_file,omitempty"`
}

func NewDynamicSource(config []byte, sourcesDir string) (Source, error) {
	var sourceRequest sourceRequest
	if err := json.NewDecoder(bytes.NewReader(config)).Decode(&sourceRequest); err != nil {
		return Source{}, fmt.Errorf("Invalid dynamic source config: %s", err)
	}

	if sourceRequest.Params.SourceFile != "" {
		source, err := ioutil.ReadFile(filepath.Join(sourcesDir, sourceRequest.Params.SourceFile))
		if err != nil {
			return Source{}, fmt.Errorf("Invalid dynamic source config: %s", err)
		}

		var tempSource Source
		yaml.Unmarshal(source, &tempSource)

		jsonBytes, err := json.Marshal(tempSource)
		if err != nil {
			return Source{}, fmt.Errorf("Invalid dynamic source config: %s", err)
		}

		if err := json.Unmarshal(jsonBytes, &sourceRequest.Source); err != nil {
			return Source{}, fmt.Errorf("Invalid dynamic source config: %s", err)
		}
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
	if source.Username == "" {
		missingParameters = append(missingParameters, "username")
	}
	if source.Password == "" {
		missingParameters = append(missingParameters, "password")
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
