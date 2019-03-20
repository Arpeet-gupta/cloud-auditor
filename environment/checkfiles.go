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

	credentialsExists, credentialsError := isCredentialsPresent(homeDir)
	if credentialsError != nil {
		config.Logger.Error(credentialsError.Error())
	}

	profile := config.Profile
	if configAWSExists {
		profilesInConfig := getProfilesFromFile(config, homeDir+"/.aws/config")
		if !helpers.SliceContains(profilesInConfig, profile) {
			var ans string
			config.Logger.GetInput("You don't have the "+profile+" profile in your config file. Would you like to create one? *y* / *n*", &ans)
			if strings.ToUpper(ans) == "y" {
				region := getUserRegion(config)
				CreateAWSConfigFile(config, profile, region, "")
			} else {
				config.Logger.Info("You can use another profile by setting the \"-p\" argument or specify a different default profile by setting the AWS_PROFILE variable")
				return false
			}
		}
		if credentialsExists {
			addProfileToCredentials(profile, homeDir, config)
		} else {
			CreateAWSCredentialsFile(config, profile)
		}
	} else {
		if credentialsExists {
			var ans string
			config.Logger.GetInput("File .aws/config does not exist, but .aws/credentials has been found. Do you want to create config file using one of the profiles in the .aws/credentias? *y* / *n*", &ans)
			if strings.ToUpper(ans) == "y" {
				createConfigProfileFromCredentials(homeDir, config, profile)
				return true
			} else {
				profile = setProfileInfoAndCreateConfigFile(config)
				CreateAWSCredentialsFile(config, profile)
			}
		} else {
			config.Logger.Info("File .aws/config does not exist.")
			profile = setProfileInfoAndCreateConfigFile(config)
			CreateAWSCredentialsFile(config, profile)
		}
	}
	return true
}
