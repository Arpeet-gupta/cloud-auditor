package resource

import (
	"github.com/iamabhishek-dubey/cloud-auditor/configuration"
	"github.com/iamabhishek-dubey/cloud-auditor/session"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/ec2"
)

type Snapshots []*ec2.Snapshot

func (s *Snapshots) LoadFromAWS(config *configuration.Config, region string) error {
	ec2API, err := config.ClientFactory.GetEc2Client(csasession.SessionConfig{Profile: config.Profile, Region: region})
	if err != nil {
		return err
	}

	q := &ec2.DescribeSnapshotsInput{
		OwnerIds: []*string{aws.String("self")},
	}

	for {
		result, err := ec2API.DescribeSnapshots(q)
		if err != nil {
			if aerr, ok := err.(awserr.Error); ok {
				switch aerr.Code() {
					case "OptInRequired":
						break
					default:
						return err
				}
			} else {
				return err
			}
		}
		*s = append(*s, result.Snapshots...)
	}
}
