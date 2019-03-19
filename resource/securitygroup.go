package resource

import (
	"github.com/iamabhishek-dubey/cloud-auditor/configuration"
	"github.com/iamabhishek-dubey/cloud-auditor/session"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/ec2"
)

type SecurityGroups map[string][]*ec2.IpPermission

func (s *SecurityGroups) LoadFromAWS(config *configuration.Config, region string) error {
	
}