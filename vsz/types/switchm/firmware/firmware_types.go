package firmware

// API Version: v8_0

type AllFirmwaresQueryResultList struct {
	*FirmwaresQueryResultList
}

type FirmwaresQueryResultList struct {
	// Extra
	// Extra information for Firmware list
	Extra interface{} `json:"extra,omitempty"`

	// FirstIndex
	// Index of the first firmware list returned out of the complete Firmware list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicator of whether there are more Firmwares after the current displayed list
	HasMore *bool `json:"hasMore,omitempty"`

	List []*SwitchFirmware `json:"list,omitempty"`

	// RawDataTotalCount
	// Firmware list count
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Total Firmware list count
	TotalCount *int `json:"totalCount,omitempty"`
}

type ScheduleIds struct {
	// Extra
	// Extra information for Schedule Ids list
	Extra interface{} `json:"extra,omitempty"`

	// FirstIndex
	// Index of the first Schedule Ids returned out of the complete ConfigBackup list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicator of whether there are more Schedule Ids after the current displayed list
	HasMore *bool `json:"hasMore,omitempty"`

	List []string `json:"list,omitempty"`

	// RawDataTotalCount
	// Firmware list count
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Total Schedule Ids count
	TotalCount *int `json:"totalCount,omitempty"`
}

type SwitchFirmware struct {
	SwitchModels []*SwitchModel `json:"switchModels,omitempty"`

	// Version
	// Firmware version of the Switch
	Version *string `json:"version,omitempty"`
}

type SwitchModel struct {
	// ImageFileNames
	// Name of the Switch Image File
	ImageFileNames []string `json:"imageFileNames,omitempty"`

	// Name
	// Name of the Switch Model
	Name *string `json:"name,omitempty"`
}
