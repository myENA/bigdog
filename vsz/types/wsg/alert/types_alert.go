package alert

// API Version: v8_1

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
)

type AckBulkAlarms struct {
	IdList []string `json:"idList,omitempty"`
}

func NewAckBulkAlarms() *AckBulkAlarms {
	ackBulkAlarmsType := new(AckBulkAlarms)
	return ackBulkAlarmsType
}

func NewAckBulkAlarmsWithDefaults() *AckBulkAlarms {
	ackBulkAlarmsType := new(AckBulkAlarms)
	return ackBulkAlarmsType
}

type ClearBulkAlarms struct {
	Comment *common.NormalNameAllowBlank `json:"comment,omitempty"`

	IdList []string `json:"idList,omitempty"`
}

func NewClearBulkAlarms() *ClearBulkAlarms {
	clearBulkAlarmsType := new(ClearBulkAlarms)
	return clearBulkAlarmsType
}

func NewClearBulkAlarmsWithDefaults() *ClearBulkAlarms {
	clearBulkAlarmsType := new(ClearBulkAlarms)
	return clearBulkAlarmsType
}
