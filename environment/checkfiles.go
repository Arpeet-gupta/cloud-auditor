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

func isAWSConfigPresent(homePath string) (bool, error) {
	_, credentialsError := os.Open(homePath + "/.aws/config")
	if credentialsError != nil {
		return false, nil
	}
	return true, nil
}

func isCredentialsPresent(homePath string) (bool, error) {
	_, credentialsError := os.Open(homePath + "/.aws/credentials")
	if credentialsError != nil {
		return false, nil
	}
	return true, nil
}

func getAllRegions() (Regions []string) {
	rs, _ := endpoints.RegionsForService(endpoints.DefaultPartitions(), endpoints.AwsPartitionID, endpoints.ApigatewayServiceID)
	for region := range rs {
		Regions = append(Regions, region)
	}
	return
}

func getUserRegion(config *configuration.Config) string {
	showAvailableRegions(config)
	var numberRegion int
	config.Logger.GetInput("Region", &numberRegion)

	for numberRegion < 0 || numberRegion >= len(Regions) {
		config.Logger.Always("Try again, invalid region")
		config.Logger.GetInput("Region", &numberRegion)
	}

	region := Regions[numberRegion]
	config.Logger.Always("Your region is: " + region)
	return region
}

func showAvailableRegions(config *configuration.Config) {
	config.Logger.Always("Available Regions:")
	for i := 0; i < len(Regions); i++ {
		pom := strconv.Itoa(i)
		config.Logger.Always("Number " + pom + " region " + Regions[i])
	}
}

func getUserOutput(config *configuration.Config) string {
	var output string
	config.Logger.GetInput("Input the output format [json, text, table]", &output)
	for !helpers.SliceContains([]string{"json", "text", "table"}, output) {
		config.Logger.Always("Try again, invalid output")
		config.Logger.GetInput("Input the output format [json, text, table]", &output)
	}
	config.Logger.Always("Your output is: " + output)
	return output
}