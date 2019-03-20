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
		path := homePath + "/.aws/credentials"
		line := "\n[" + profile + "-long-term" + "]\n"
		appendStringToFile(path, line)
		line = "aws_access_key_id" + " = " + awsAccessKeyID + "\n"
		appendStringToFile(path, line)
		line = "aws_secret_access_key" + " = " + awsSecretAccessKey + "\n"
		appendStringToFile(path, line)
		line = "mfa_serial" + " = " + mfaSerial + "\n"
		appendStringToFile(path, line)
	}
}

func CreateAWSConfigFile(config *configuration.Config, profile string, region string, output string) {
	if output == "" {
		output = getUserOutput(config)
	}
	homePath, pathError := GetUserHomeDir()
	if pathError != nil {
		config.Logger.Error(pathError.Error())
	}
	path := homePath + "/.aws/config"
	line := "\n[" + profile + "]\n"
	appendStringToFile(path, line)
	line = "region" + " = " + region + "\n"
	appendStringToFile(path, line)
	line = "output" + " = " + output + "\n"
	appendStringToFile(path, line)
}

func createConfigProfileFromCredentials(homeDir string, config *configuration.Config, profile string) {
	profilesInCredentials := UniqueNonEmptyElementsOf(getProfilesFromFile(config, homeDir+"/.aws/credentials"))
	config.Logger.Always("Available profile names are: " + fmt.Sprint("[ "+strings.Join(profilesInCredentials, ", ")+" ]"))
	config.Logger.GetInput("Profile", &profile)

	for !helpers.SliceContains(profilesInCredentials, profile) {
		config.Logger.Always("Invalid profile name. Available profile names are: " + fmt.Sprint("[ "+strings.Join(profilesInCredentials, ", ")+" ]"))
		config.Logger.GetInput("Profile", &profile)
	}
	region := getUserRegion(config)
	CreateAWSConfigFile(config, profile, region, "")
}
