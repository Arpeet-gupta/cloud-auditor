package configuration

import (
	"github.com/iamabhishek-dubey/cloud-auditor/csasession/sessionfactory"
	"github.com/iamabhishek-dubey/cloud-auditor/csasession/clientfactory"
	"github.com/iamabhishek-dubey/cloud-auditor/logger"
)

type Config struct {
	Regions        *[]string
	Services       *[]string
	Profile        string
	SessionFactory *sessionfactory.SessionFactory
	ClientFactory  clientfactory.ClientFactory
	Logger         *logger.Logger
}

func GetConfig() (config Config) {
	myLogger := logger.CreateDefaultLogger()
	config.Logger = &myLogger
	config.SessionFactory = sessionfactory.New()
	config.ClientFactory = clientfactory.New(config.SessionFactory)

	return config
}

