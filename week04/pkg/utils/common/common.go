package common

import (
	"golang.org/x/tools/go/ssa/interp/testdata/src/errors"
	"strings"
)

//RepairPatientId patient-id Standardization, eg: A12345 -> 0000A12345
func RepairPatientId(pid string) (string, error) {
	rawLen := len(pid)
	if rawLen > 10 {
		return "", errors.New("invalid patient id, the patient-id can not exceed 10 digits")
	}
	var sb strings.Builder
	for i := 0; i < 10-rawLen; i++ {
		sb.WriteString("0")
	}
	sb.WriteString(pid)
	return sb.String(), nil
}
