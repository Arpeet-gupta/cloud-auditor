package csasession

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

// SessionConfig for providing configs of AWS
type SessionConfig struct {
	Profile string
	Region string
}

// CreateSession returns new AWS Session
func CreateSession(config SessionConfig) (*session.Session, error) {
	sess, err := session.NewSessionWithOptions( 
		session.Options{
			Config: aws.Config{
				Region: &config.Region,
			},
			Profile: config.Profile,
		})
	return sess, err
}

// GetAvailableRegions returns list of All AWS regions
func GetAvailableRegions() *[]string {
	return &[]string{
		"us-east-2",
		"us-east-1",
		"us-west-1",
		"us-west-2",
	}
}
