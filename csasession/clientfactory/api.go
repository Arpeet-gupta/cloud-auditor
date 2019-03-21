package clientfactory

import (
	"github.com/iamabhishek-dubey/cloud-auditor/csasession"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/kms"
)

type ClientFactory interface {
	GetEc2Client(config csasession.SessionConfig) (EC2Client, error)
	GetKmsClient(config csasession.SessionConfig) (KmsClient, error)
}

// GetKmsClient creates a new KMS client from cached session.
func (factory *ClientFactoryAWS) GetKmsClient(config csasession.SessionConfig) (KmsClient, error) {
	sess, err := factory.sessionFactory.GetSession(config)
	if err != nil {
		return nil, err
	}

	client := kms.New(sess)
	return AWSKmsClient{api: client}, nil
}

// GetEc2Client creates a new EC2 client from cached session.
func (factory *ClientFactoryAWS) GetEc2Client(config csasession.SessionConfig) (EC2Client, error) {
	sess, err := factory.sessionFactory.GetSession(config)
	if err != nil {
		return nil, err
	}

	client := ec2.New(sess)
	return AWSEC2Client{api: client}, nil
}

