package environment

import (
	"bufio"
	"github.com/iamabhishek-dubey/cloud-auditor/configuration"
	"github.com/iamabhishek-dubey/cloud-auditor/helpers"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"os"
	"strconv"
	"strings"
)

var Regions = getAllRegions()

func CheckAWSConfigFiles(config *configuration.Config) bool {
	homeDir, pathError := GetUserHomeDir()
	if pathError != nil {
		config.Logger.Error(pathError.Error())
		return false
	}
	
	configAWSExists, configError := isAWSConfigPresent(homeDir)
	if configError != nil {
		config.Logger.Error(configError.Error())
	}
}