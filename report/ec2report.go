package report

import (
	"fmt"
	"bytes"
	"github.com/iamabhishek-dubey/cloud-auditor/configuration"
	"github.com/iamabhishek-dubey/cloud-auditor/environment"
	"github.com/iamabhishek-dubey/cloud-auditor/resource"
	"sort"
	"strconv"
	"strings"
	"os"
	"html/template"
)

var (
	azs [][]string
)

type Ec2Report struct {
	VolumeReport      *VolumeReport
	InstanceID        string
	SortableTags      *SortableTags
	SecurityGroupsIDs []string
	AvailabilityZone  string
}

type PageData struct {
	PageTitle string
	Avz       []string
}

func NewEc2Report(instanceID string) *Ec2Report {
	return &Ec2Report{
		InstanceID:   instanceID,
		VolumeReport: &VolumeReport{},
		SortableTags: NewSortableTags(),
	}
}

type Ec2Reports []*Ec2Report

type Ec2ReportRequiredResources struct {
	Ec2s             *resource.Ec2s
	KMSKeys          *resource.KMSKeys
	Volumes          *resource.Volumes
	SecurityGroups   *resource.SecurityGroups
	AvailabilityZone string
}

func (e *Ec2Reports) GetHeaders() []string {
	return []string{"Availability\nZone", "EC2", "Volumes\n(None) - not encrypted\n(DKMS) - encrypted with default KMSKey", "Security\nGroups\n(Incoming CIDR = 0\x2E0\x2E0\x2E0/0)\nID : PROTOCOL : PORT", "EC2 Tags"}
}

func (e *Ec2Reports) FormatDataToTable() [][]string {
	data := [][]string{}
	datas := [][]string{}

	lsj := make(map[string][][]string)
	for _, ec2Report := range *e {
		row := []string{
			ec2Report.AvailabilityZone,
			ec2Report.InstanceID,
			ec2Report.VolumeReport.ToTableData(),
			SliceOfStringsToString(ec2Report.SecurityGroupsIDs),
			ec2Report.SortableTags.ToTableData(),
		}
//	for _, ec2Report := range *e {
		rows := []string{
		    ec2Report.AvailabilityZone,
		}
		datas = append(datas, rows)
		fmt.Println(datas)
//	 }
		//fmt.Println(azs)
//			datas := PageData{
//				PageTitle: "Test Page",
//				Avz: []string{
//					azs,
//				},
//			}

//			tmpl := template.Must(template.ParseFiles("view/layout.html"))
//			tmpl.Execute(os.Stdout, datas)
		data = append(data, row)
		//fmt.Println(data)
		//datas = append(datas, rows)
		//fmt.Println(datas)
	}
	srtdata := sortTableData(datas)
//	lsdata := srtdata.(map[string][][]string)
        lsj["t"] = srtdata
	azs = lsj["t"]
	fmt.Println(lsj)
//	lst :=  string(srtdata)
//	fmt.Println(lst)
	htdata := PageData{
	    PageTitle: "Hey",
	    Avz: []string{
	       "hi",
	    },
	}
	tmpl := template.Must(template.ParseFiles("view/layout.html"))
	tmpl.Execute(os.Stdout, htdata)
	sortedData := sortTableData(data)
	return sortedData
}

func (e *Ec2Reports) GenerateReport(r *Ec2ReportRequiredResources) {
	for _, ec2 := range *r.Ec2s {
		ec2Report := NewEc2Report(*ec2.InstanceId)
		ec2OK := true
		for _, blockDeviceMapping := range ec2.BlockDeviceMappings {
			volume := r.Volumes.FindById(*blockDeviceMapping.Ebs.VolumeId)
			if !*volume.Encrypted {
				ec2OK = false
				ec2Report.VolumeReport.AddEBS(*volume.VolumeId, NONE)
			} else {
				kmskey := r.KMSKeys.FindByKeyArn(*volume.KmsKeyId)
				if !kmskey.Custom {
					ec2OK = false
					ec2Report.VolumeReport.AddEBS(*volume.VolumeId, DKMS)
				}
			}
		}

		for _, sg := range ec2.SecurityGroups {
			ipPermissions := r.SecurityGroups.GetIpPermissionsByID(*sg.GroupId)
			if ipPermissions != nil {
				for _, ipPermission := range ipPermissions {
					for _, ipRange := range ipPermission.IpRanges {
						if *ipRange.CidrIp == "0.0.0.0/0" {
							ec2Report.SecurityGroupsIDs = append(ec2Report.SecurityGroupsIDs, *sg.GroupId+" : "+*ipPermission.IpProtocol+" : "+strconv.FormatInt(*ipPermission.ToPort, 10))
							ec2OK = false
						}
					}
				}
			}
		}
		if !ec2OK {
			ec2Report.SortableTags.Add(ec2.Tags)
			*e = append(*e, ec2Report)
		}
		ec2Report.AvailabilityZone = *ec2.Placement.AvailabilityZone
		//azs = *ec2.Placement.AvailabilityZone
//		data := PageData{
//			PageTitle: "Test Page",
//			Avz: ec2Report.AvailabilityZone,
//		}

		//tmpl := template.Must(template.ParseFiles("view/layout.html"))
		//tmpl.Execute(os.Stdout, data)
	}
//	fmt.Println(azs)
}

// GetResources : Initialize and loads required resources to create ec2 report
func (e *Ec2Reports) GetResources(config *configuration.Config) (*Ec2ReportRequiredResources, error) {
	resources := &Ec2ReportRequiredResources{
		KMSKeys:          resource.NewKMSKeys(),
		Ec2s:             &resource.Ec2s{},
		Volumes:          &resource.Volumes{},
		SecurityGroups:   &resource.SecurityGroups{},
		AvailabilityZone: "zone",
	}

	for _, region := range *config.Regions {
		err := resource.LoadResources(
			config,
			region,
			resources.Ec2s,
			resources.KMSKeys,
			resources.Volumes,
			resources.SecurityGroups,
		)
		if err != nil {
			return nil, err
		}
	}
	return resources, nil
}

func SliceOfStringsToString(slice []string) string {
	n := len(slice)
	if n == 0 {
		return ""
	}
	var buffer bytes.Buffer
	for _, s := range slice[:n-1] {
		buffer.WriteString(s + "\n")
	}
	buffer.WriteString(slice[n-1])
	return buffer.String()
}

func sortTableData(data [][]string) [][]string {
	if data[0][0] == "" {
		return data
	}
	var regions []string
	var sortedData [][]string

	for _, regs := range data {
		reg := regs[0][:len(regs[0])-1]
		regions = append(regions, reg)
	}
	sort.Strings(regions)
	uniqueregions := environment.UniqueNonEmptyElementsOf(regions)
	for _, unique := range uniqueregions {
		for _, b := range data {
			if strings.Contains(b[0], unique) {
				sortedData = append(sortedData, b)
			}
		}
	}
	return sortedData
}
