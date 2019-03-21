package scanner

import (
	"fmt"
	"strings"
	"github.com/iamabhishek-dubey/cloud-auditor/configuration"
	"github.com/iamabhishek-dubey/cloud-auditor/report"
)

func Run(config *configuration.Config) error {
	for _, service := range *config.Services {
		switch strings.ToLower(service) {
			
		}
	}
}

