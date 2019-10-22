package alert

// API Version: v8_1

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
)

type AckBulkAlarms struct {
	IdList []string `json:"idList,omitempty"`
}

type ClearBulkAlarms struct {
	Comment *common.NormalNameAllowBlank `json:"comment,omitempty"`

	IdList []string `json:"idList,omitempty"`
}

