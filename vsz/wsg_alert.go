package vsz

// API Version: v9_0

type WSGAlertAckBulkAlarms struct {
	// IdList
	// Constraints:
	//    - nullable
	IdList []string `json:"idList,omitempty" validate:"omitempty,dive"`
}

func NewWSGAlertAckBulkAlarms() *WSGAlertAckBulkAlarms {
	m := new(WSGAlertAckBulkAlarms)
	return m
}

type WSGAlertClearBulkAlarms struct {
	// Comment
	// Constraints:
	//    - nullable
	Comment *WSGCommonNormalNameAllowBlank `json:"comment,omitempty"`

	// IdList
	// Constraints:
	//    - nullable
	IdList []string `json:"idList,omitempty" validate:"omitempty,dive"`
}

func NewWSGAlertClearBulkAlarms() *WSGAlertClearBulkAlarms {
	m := new(WSGAlertClearBulkAlarms)
	return m
}
