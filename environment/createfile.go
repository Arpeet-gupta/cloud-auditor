package environment

import (
	"fmt"
	"github.com/iamabhishek-dubey/cloud-auditor/configuration"
	"github.com/iamabhishek-dubey/cloud-auditor/helpers"
	"os"
	"strings"
)

func CreateAWSCredentialsFile(config *configuration.Config, profile string) {
	if profile != "" {
		config.Logger.Always("You haven't got .aws/credentials file for profile " + profile)
		var awsAccessKeyID string
		var awsSecretAccessKey string

		config.Logger.GetInput("awsAccessKeyID", &awsAccessKeyID)
		config.Logger.GetInput("awsSecretAccessKey", &awsSecretAccessKey)

		homePath, pathError := GetUserHomeDir()
		
		if pathError != nil {
			config.Logger.Error(pathError.Error())
		}
	}
}
