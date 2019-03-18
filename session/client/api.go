package client

import (
	"github.com/iamabhishek-dubey/cloud-auditor/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

type ClientFactory interface {
	GetEc2Client(config csasession.SessionConfig) (EC2Client, error)
}

func (factory *ClientFactoryAWS) GetEc2Client(config csasession.SessionConfig) (EC2Client, error {
	sess, err := factory.sessionFactory.GetSession(config)
	if err != nil {
		return nil, err
	}

	client := ec2.New(sess)
	return AWSEC2Client{api: client}, nil
}
