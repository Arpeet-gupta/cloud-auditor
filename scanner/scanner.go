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
		case "ec2":
			config.Logger.Info("Gathering information about EC2s...")
			ec2Reports := report.Ec2Reports{}
			resources, err := ec2Reports.GetResources(config)
			if err != nil {
				return err
			}
			ec2Reports.GenerateReport(resources)
			test := ec2Reports.GenerateReport(resources)
			fmt.Println(test)
			report.PrintTable(&ec2Reports)

		default:
			return fmt.Errorf("Wrong service name: %s", service)
		}
	}
	return nil
}

