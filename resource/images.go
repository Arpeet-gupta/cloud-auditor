package resource

import (
	"github.com/iamabhishek-dubey/cloud-auditor/configuration"
	"github.com/iamabhishek-dubey/cloud-auditor/session"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/ec2"
)

type Images []*ec2.Image

func (im Images) SortByDate() {
	sort.SliceStable(im, func(i, j int) bool {
		return *(im)[i].CreationDate < *(im)[j].CreationDate
	}
}

func (im *Images) FindByTags(tags map[string]string) Images {
	found := Images{}
	n := 0
	for _, image := range *im {

	}
}
