package vsz

// API Version: v9_0

type WSGAlertAckBulkAlarms struct {
	IdList []string `json:"idList"`
}

func NewWSGAlertAckBulkAlarms() *WSGAlertAckBulkAlarms {
	m := new(WSGAlertAckBulkAlarms)
	return m
}

type WSGAlertClearBulkAlarms struct {
	Comment *WSGCommonNormalNameAllowBlank `json:"comment,omitempty"`

	IdList []string `json:"idList"`
}

func NewWSGAlertClearBulkAlarms() *WSGAlertClearBulkAlarms {
	m := new(WSGAlertClearBulkAlarms)
	return m
}
