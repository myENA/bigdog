package alert

// API Version: v8_0

type AckBulkAlarms struct {
	IDList []string `json:"idList,omitempty"`
}

type ClearBulkAlarms struct {
	Comment *string  `json:"comment,omitempty"`
	IDList  []string `json:"idList,omitempty"`
}
