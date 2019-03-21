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

func PrintTable(r Report) {
	data := r.FormatDataToTable()

	table := tablewriter.NewWriter(os.Stdout)
	//Headers
	table.SetReflowDuringAutoWrap(false)
	table.SetAutoFormatHeaders(false)
	table.SetHeader(customFormatHeaders(r.GetHeaders()))
	//Configure Rows and Cells
	table.SetRowSeparator("-")
	table.SetRowLine(true)
	table.SetAutoWrapText(false)
	table.AppendBulk(data)
	table.Render()
}

func customFormatHeaders(headers []string) []string {
	for i, header := range headers {
		headers[i] = Title(header)
	}
	return headers
}

func Title(name string) string {
	origLen := len(name)
	name = strings.Replace(name, "_", " ", -1)
	//name = strings.Replace(name, ".", " ", -1)
	name = strings.TrimSpace(name)
	if len(name) == 0 && origLen > 0 {
		// Keep at least one character. This is important to preserve
		// empty lines in multi-line headers/footers.
		name = " "
	}
	return strings.ToUpper(name)
}

type EncryptionType int
