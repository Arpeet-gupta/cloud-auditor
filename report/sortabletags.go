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
