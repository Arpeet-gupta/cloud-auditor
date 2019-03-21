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
