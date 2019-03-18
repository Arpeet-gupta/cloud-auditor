package session

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

type SessionConfig struct {
	Profile string
	Region string
}

func CreateSession(config SessionConfig) (*session.Session, error) {
	sess, err := session.NewSessionWithOptions( 
		session.Options{
			Config: aws.Config{
				Region: &config.Region,
			},
			Profile: config.Profile,
		}
	)
	return sess, err
}

func GetAvailableRegions() *[]string {
	return &[]string{
		"us-east-2",
		"us-east-1",
		"us-west-1",
		"us-west-2"
	}
}
