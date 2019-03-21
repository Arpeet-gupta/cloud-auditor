package report

import (
	"bytes"
	"sort"
	"strings"
	"github.com/aws/aws-sdk-go/service/ec2"
)

type SortableTags struct {
	Tags map[string]string
	Keys []string
}

func NewSortableTags() *SortableTags {
	return &SortableTags{Tags: make(map[string]string)}
}

func (st *SortableTags) Add(tags []*ec2.Tag) {
	for _, tag := range tags {
		st.Keys = append(st.Keys, *tag.Key)
		st.Tags[*tag.Key] = *tag.Value
	}
	less := func(i, j int) bool {
		return strings.ToLower(st.Keys[i]) < strings.ToLower(st.Keys[j])
	}
	sort.Slice(st.Keys, less)
}
