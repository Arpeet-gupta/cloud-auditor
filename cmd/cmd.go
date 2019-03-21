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
			
		}
	}
}