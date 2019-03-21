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
