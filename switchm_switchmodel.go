package bigdog

// API Version: v9_1

// SwitchMSwitchModelModelList
//
// Definition: switchModel_modelList
type SwitchMSwitchModelModelList struct {
	// Model
	// Switch Model
	Model *string `json:"model,omitempty"`
}

func NewSwitchMSwitchModelModelList() *SwitchMSwitchModelModelList {
	m := new(SwitchMSwitchModelModelList)
	return m
}

// SwitchMSwitchModelResult
//
// Definition: switchModel_result
type SwitchMSwitchModelResult struct {
	// Extra
	// Extra field
	Extra interface{} `json:"extra,omitempty"`

	// FirstIndex
	// The first data index for current reulst
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicator of remaining data
	HasMore *bool `json:"hasMore,omitempty"`

	List []*SwitchMSwitchModelModelList `json:"list,omitempty"`

	// RawDataTotalCount
	// Total Data Count
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Total Data Count
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewSwitchMSwitchModelResult() *SwitchMSwitchModelResult {
	m := new(SwitchMSwitchModelResult)
	return m
}
