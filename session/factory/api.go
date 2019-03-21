package sessionfactory

import (
	"github.com/iamabhishek-dubey/cloud-auditor/session"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// GetSession returns cached AWS session. For faster AWS Config
func (factory *SessionFactory) GetSession(config session.SessionConfig) (*session.Session, error) {
	factory.mutex.Lock()
	defer factory.mutex.Unlock()
	if sess, ok := factory.regionToSession[config.Region]; ok {
		return sess, nil
	}
	return factory.NewSession(config)
}

// NewSession creates a new session and caches it.
func (factory *SessionFactory) NewSession(config session.SessionConfig) (*session.Session, error) {
	sess, err := session.CreateSession(config)
	if err != nil {
		return nil, err
	}

	factory.regionToSession[config.Region] = sess
	return sess, nil
}

func (factory *SessionFactory) SetNormalizeBucketLocation(config session.SessionConfig) error {
	sess, err := factory.GetSession(config)
	if err != nil {
		return err
	}
	sess.Handlers.Unmarshal.PushBackNamed(s3.NormalizeBucketLocationHandler)
	return nil
}

func (factory *SessionFactory) ReinitialiseSession(config session.SessionConfig) (err error) {
	factory.mutex.Lock()
	defer factory.mutex.Unlock()

	_, err = factory.NewSession(config)
	return
}
