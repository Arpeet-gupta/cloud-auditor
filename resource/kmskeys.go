package resource

import (
	"fmt"
	"strings"
	"sync"
	"github.com/iamabhishek-dubey/cloud-auditor/configuration"
	"github.com/iamabhishek-dubey/cloud-auditor/session"
	"github.com/iamabhishek-dubey/cloud-auditor/session/client"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/kms"
)

type KMSKey struct {
	AliasArn  string
	AliasName string
	Custom    bool
	KeyId     string // the same as TargetKeyId in AliasListEntry
}

type KMSKeys struct {
	Values map[string]*KMSKey
	sync.RWMutex
}

// NewKMSKeys : Initialize KMS Keys struct with map of keys
func NewKMSKeys() *KMSKeys {
	return &KMSKeys{Values: make(map[string]*KMSKey)}
}

type KMSKeyAliases []*kms.AliasListEntry

type KMSKeysListEntries []*kms.KeyListEntry

// LoadAllFromAWS : Load KMS Keys from all regions
func (k *KMSKeys) LoadAllFromAWS(config *configuration.Config) error {
	regions := *csasession.GetAvailableRegions()

	var wg sync.WaitGroup
	n := len(regions) * 2
	done := make(chan bool, n)
	errc := make(chan error, n)
	wg.Add(n)

	go func() {
		wg.Wait()
		close(done)
		close(errc)
	}()

	kmsKeyAliases := &KMSKeyAliases{}
	kmsKeyListEntries := &KMSKeysListEntries{}

	for _, region := range regions {
		kmsClient, err := config.ClientFactory.GetKmsClient(csasession.SessionConfig{Profile: config.Profile, Region: region})
		if err != nil {
			return err
		}

		go loadKeyListEntries(kmsClient, kmsKeyListEntries, done, errc, &wg)
		go loadKeyAliases(kmsClient, kmsKeyAliases, done, errc, &wg)
	}
	for i := 0; i < n; i++ {
		select {
			case <-done:
			case err := <-errc:
				return err
		}
	}
	k.loadValuesToMap(kmsKeyAliases, kmsKeyListEntries)
	return nil
}
