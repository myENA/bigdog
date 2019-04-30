package firmware

// API Version: v8_0

type AllFirmwaresQueryResultList struct {
	*FirmwaresQueryResultList
}

type FirmwaresQueryResultList struct {
	Extra             interface{}       `json:"extra,omitempty"`
	FirstIndex        *int              `json:"firstIndex,omitempty"`
	HasMore           *bool             `json:"hasMore,omitempty"`
	List              []*SwitchFirmware `json:"list,omitempty"`
	RawDataTotalCount *int              `json:"rawDataTotalCount,omitempty"`
	TotalCount        *int              `json:"totalCount,omitempty"`
}

type ScheduleIds struct {
	Extra             interface{} `json:"extra,omitempty"`
	FirstIndex        *int        `json:"firstIndex,omitempty"`
	HasMore           *bool       `json:"hasMore,omitempty"`
	List              []string    `json:"list,omitempty"`
	RawDataTotalCount *int        `json:"rawDataTotalCount,omitempty"`
	TotalCount        *int        `json:"totalCount,omitempty"`
}

type SwitchFirmware struct {
	SwitchModels []*SwitchModel `json:"switchModels,omitempty"`
	Version      *string        `json:"version,omitempty"`
}

type SwitchModel struct {
	ImageFileNames []string `json:"imageFileNames,omitempty"`
	Name           *string  `json:"name,omitempty"`
}
