package alert

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/wsg/common"
)

type AckBulkAlarms struct {
	IDList []string `json:"idList,omitempty"`
}

type ClearBulkAlarms struct {
	Comment *common.NormalNameAllowBlank `json:"comment,omitempty"`

	IDList []string `json:"idList,omitempty"`
}
