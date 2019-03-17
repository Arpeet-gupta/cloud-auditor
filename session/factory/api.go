package sessionfactory

import (
	"github.com/iamabhishek-dubey/cloud-auditor/session"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func (factory *SessionFactory) GetSession(config csasession.SessionConfig) (*session.Session, error) {
	factory.mutex.Lock()
	defer factory.mutex.Unlock()
	if sess, ok := factory.regionToSession[config.Region]; ok {
		return sess, nil
	}
	return factory.NewSession(config)
}

func (factory *SessionFactory) NewSession(config csasession.SessionConfig) (*session.Session, error) {
	sess, err := csasession.CreateSession(config)
	if err != nil {
		return nil, err
	}

	factory.regionToSession[config.Region] = sess
	return sess, nil
}

func (factory *SessionFactory) SetNormalizeBucketLocation(config csasession.SessionConfig) error {
	sess, err := factory.GetSession(config)
	if err != nil {
		return err
	}
	sess.Handlers.Unmarshal.PushBackNamed(s3.NormalizeBucketLocationHandler)
	return nil
}

func (factory *SessionFactory) ReinitialiseSession(config csasession.SessionConfig) (err error) {
	factory.mutex.Lock()
	defer factory.mutex.Unlock()

	_, err = factory.NewSession(config)
	return
}
