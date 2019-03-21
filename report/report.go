package report

import (
	"github.com/olekukonko/tablewriter"
	"os"
	"strings"
)

type Report interface {
	FormatDataToTable() [][]string
	GetHeaders() []string
}
