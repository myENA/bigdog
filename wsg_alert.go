package bigdog

// API Version: v9_1

// WSGAlertAckBulkAlarms
//
// Definition: alert_ackBulkAlarms
type WSGAlertAckBulkAlarms struct {
	IdList []string `json:"idList,omitempty"`
}

func NewWSGAlertAckBulkAlarms() *WSGAlertAckBulkAlarms {
	m := new(WSGAlertAckBulkAlarms)
	return m
}

// WSGAlertClearBulkAlarms
//
// Definition: alert_clearBulkAlarms
type WSGAlertClearBulkAlarms struct {
	Comment *WSGCommonNormalNameAllowBlank `json:"comment,omitempty"`

	IdList []string `json:"idList,omitempty"`
}

func NewWSGAlertClearBulkAlarms() *WSGAlertClearBulkAlarms {
	m := new(WSGAlertClearBulkAlarms)
	return m
}
