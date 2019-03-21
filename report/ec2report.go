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

func (e *Ec2Reports) FormatDataToTable() [][]string {
	data := [][]string{}

	for _, ec2Report := range *e {
		row := []string{
			ec2Report.AvailabilityZone,
			ec2Report.InstanceID,
			ec2Report.VolumeReport.ToTableData(),
			SliceOfStringsToString(ec2Report.SecurityGroupsIDs),
			ec2Report.SortableTags.ToTableData(),
		}
		data = append(data, row)
	}
	sortedData := sortTableData(data)
	return sortedData
}

func (e *Ec2Reports) GenerateReport(r *Ec2ReportRequiredResources) {
	for _, ec2 := range *r.Ec2s {
		ec2Report := NewEc2Report(*ec2.InstanceId)
		ec2OK := true
		for _, blockDeviceMapping := range ec2.BlockDeviceMappings {
			
		}
	}
}