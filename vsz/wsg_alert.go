package vsz

// API Version: v8_1

type WSGAlertAckBulkAlarms struct {
	IdList []string `json:"idList,omitempty"`
}

func NewWSGAlertAckBulkAlarms() *WSGAlertAckBulkAlarms {
	m := new(WSGAlertAckBulkAlarms)
	return m
}

type WSGAlertClearBulkAlarms struct {
	Comment *WSGCommonNormalNameAllowBlank `json:"comment,omitempty"`

	IdList []string `json:"idList,omitempty"`
}

func NewWSGAlertClearBulkAlarms() *WSGAlertClearBulkAlarms {
	m := new(WSGAlertClearBulkAlarms)
	return m
}
