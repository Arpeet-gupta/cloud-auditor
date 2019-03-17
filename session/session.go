package session

import (
	"sync"
	"github.com/aws/aws-sdk-go/session"
)

type Session struct {
	regionToSession map[string]*session.Session
	mutex           sync.Mutex
}

func New() *Session {
	factory := &SessionFactory {
		regionToSession: make(map[string]*session.Session),
	}
	return factory
}
