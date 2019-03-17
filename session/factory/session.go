package sessionfactory

import (
	"sync"
	"github.com/aws/aws-sdk-go/session"
)

type SessionFactory struct {
	regionToSession map[string]*session.Session
	mutex           sync.Mutex
}

func New() *SessionFactory {
	factory := &SessionFactory {
		regionToSession: make(map[string]*session.Session),
	}
	return factory
}
