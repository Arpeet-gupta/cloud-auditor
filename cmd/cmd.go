package cmd

import (
	"os"
	"github.com/iamabhishek-dubey/cloud-auditor/configuration"
	"github.com/iamabhishek-dubey/cloud-auditor/session"
	"github.com/iamabhishek-dubey/cloud-auditor/environment"
	"github.com/iamabhishek-dubey/cloud-auditor/resource"
	"github.com/iamabhishek-dubey/cloud-auditor/scanner"
	"github.com/spf13/cobra"
)

// var cfgFile string
var config = configuration.GetConfig()

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cloud-auditor",
	Short: "Scan for vulnerabilities in your AWS Account.",
	Long:  `Scan for vulnerabilities in your AWS Account.`,
	Run: func(cmd *cobra.Command, args []string) {
		_, ok := os.LookupEnv("AWS_LAMBDA_FUNCTION_NAME") // If csa is running on lambda then env will be available. In other case csa needs config files.
		if !ok {
			if environment.CheckAWSConfigFiles(&config) {
				err := scanner.Run(&config)
				if err != nil {
					config.Logger.Error(err.Error())
				}
			}
		} else {
			err := scanner.Run(&config)
			if err != nil {
				config.Logger.Error(err.Error())
			}
		}
	},
}

// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		config.Logger.Error(err.Error())
		os.Exit(1)
	}
}

var (
	region      string
	service     string
	profile     string
)

func init() {
	
}