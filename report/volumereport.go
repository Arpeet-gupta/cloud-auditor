package report

import (
	"bytes"
	"fmt"
)

type VolumeReport []string

func (v *VolumeReport) AddEBS(volumeID string, encryptionType EncryptionType) {
	*v = append(*v, volumeID+fmt.Sprintf("[%s]", encryptionType.String()))
}

func (v *VolumeReport) ToTableData() string {
	if len(*v) == 0 {
		return ""
	}
	
}
