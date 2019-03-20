package resource

import (
	"encoding/json"
	"io/ioutil"
	"sync"
	"github.com/iamabhishek-dubey/cloud-auditor/configuration"
)

type Resource interface {
	LoadFromAWS(config *configuration.Config, region string) error
}

func LoadResource(r Resource, config *configuration.Config, region string) error {
	err := r.LoadFromAWS(config, region)
	if err != nil {
		return err
	}
	return nil
}

func LoadResources(config *configuration.Config, region string, resources ...Resource) error {
	err := r.LoadFromAWS(config, region)
	if err != nil {
		return err
	}

	return nil
}
