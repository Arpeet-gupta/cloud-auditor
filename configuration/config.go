package configuration

import (
	"github.com/iamabhishek-dubey/cloud-auditor/session/factory"
	"github.com/iamabhishek-dubey/cloud-auditor/session/client"
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
