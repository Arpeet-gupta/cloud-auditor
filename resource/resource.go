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
