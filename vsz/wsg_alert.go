package vsz

// API Version: v8_1

type WSGAlertAckBulkAlarms struct {
	IdList []string `json:"idList,omitempty"`
}

type WSGAlertClearBulkAlarms struct {
	Comment *WSGCommonNormalNameAllowBlank `json:"comment,omitempty"`

	IdList []string `json:"idList,omitempty"`
}
