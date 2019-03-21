package sessionfactory

import (
	"sync"
	"github.com/aws/aws-sdk-go/aws/session"
)

// SessionFactory provides method for creation of service client. For ex:- EC2
type SessionFactory struct {
	regionToSession map[string]*session.Session
	mutex           sync.Mutex
}

// New Creates a new session of the SessionFactory depending on region.
func New() *SessionFactory {
	factory := &SessionFactory {
		regionToSession: make(map[string]*session.Session),
	}
	return factory
}
