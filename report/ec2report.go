package report

import (
	"bytes"
	"github.com/iamabhishek-dubey/cloud-auditor/configuration"
	"github.com/iamabhishek-dubey/cloud-auditor/environment"
	"github.com/iamabhishek-dubey/cloud-auditor/resource"
	"sort"
	"strconv"
	"strings"
)

type Ec2Report struct {
	VolumeReport      *VolumeReport
	InstanceID        string
	SortableTags      *SortableTags
	SecurityGroupsIDs []string
	AvailabilityZone  string
}

func (e *Ec2Reports) GetHeaders() []string {
	return []string{"Availability\nZone", "EC2", "Volumes\n(None) - not encrypted\n(DKMS) - encrypted with default KMSKey", "Security\nGroups\n(Incoming CIDR = 0\x2E0\x2E0\x2E0/0)\nID : PROTOCOL : PORT", "EC2 Tags"}
}