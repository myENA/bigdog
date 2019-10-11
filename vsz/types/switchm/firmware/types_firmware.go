package firmware

// API Version: v8_1

import (
	"encoding/json"
)

type AllFirmwaresQueryResultList struct {
	*FirmwaresQueryResultList
}

type FirmwaresQueryResultList struct {
	// Extra
	// Extra information for Firmware list
	Extra *FirmwaresQueryResultListExtraType `json:"extra,omitempty"`

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

func NewFirmwaresQueryResultList() *FirmwaresQueryResultList {
	firmwaresQueryResultListType := new(FirmwaresQueryResultList)
	return firmwaresQueryResultListType
}

func NewFirmwaresQueryResultListWithDefaults() *FirmwaresQueryResultList {
	firmwaresQueryResultListType := new(FirmwaresQueryResultList)
	return firmwaresQueryResultListType
}

// FirmwaresQueryResultListExtraType
//
// Extra information for Firmware list
type FirmwaresQueryResultListExtraType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *FirmwaresQueryResultListExtraType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = FirmwaresQueryResultListExtraType{XAdditionalProperties: tmp}
	return nil
}

func (t *FirmwaresQueryResultListExtraType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewFirmwaresQueryResultListExtraType() *FirmwaresQueryResultListExtraType {
	firmwaresQueryResultListExtraTypeType := new(FirmwaresQueryResultListExtraType)
	return firmwaresQueryResultListExtraTypeType
}

func NewFirmwaresQueryResultListExtraTypeWithDefaults() *FirmwaresQueryResultListExtraType {
	firmwaresQueryResultListExtraTypeType := new(FirmwaresQueryResultListExtraType)
	return firmwaresQueryResultListExtraTypeType
}

type ScheduleIds struct {
	// Extra
	// Extra information for Schedule Ids list
	Extra *ScheduleIdsExtraType `json:"extra,omitempty"`

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

func NewScheduleIds() *ScheduleIds {
	scheduleIdsType := new(ScheduleIds)
	return scheduleIdsType
}

func NewScheduleIdsWithDefaults() *ScheduleIds {
	scheduleIdsType := new(ScheduleIds)
	return scheduleIdsType
}

// ScheduleIdsExtraType
//
// Extra information for Schedule Ids list
type ScheduleIdsExtraType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *ScheduleIdsExtraType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = ScheduleIdsExtraType{XAdditionalProperties: tmp}
	return nil
}

func (t *ScheduleIdsExtraType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewScheduleIdsExtraType() *ScheduleIdsExtraType {
	scheduleIdsExtraTypeType := new(ScheduleIdsExtraType)
	return scheduleIdsExtraTypeType
}

func NewScheduleIdsExtraTypeWithDefaults() *ScheduleIdsExtraType {
	scheduleIdsExtraTypeType := new(ScheduleIdsExtraType)
	return scheduleIdsExtraTypeType
}

type SwitchFirmware struct {
	SwitchModels []*SwitchModel `json:"switchModels,omitempty"`

	// Version
	// Firmware version of the Switch
	Version *string `json:"version,omitempty"`
}

func NewSwitchFirmware() *SwitchFirmware {
	switchFirmwareType := new(SwitchFirmware)
	return switchFirmwareType
}

func NewSwitchFirmwareWithDefaults() *SwitchFirmware {
	switchFirmwareType := new(SwitchFirmware)
	return switchFirmwareType
}

type SwitchModel struct {
	// ImageFileNames
	// Name of the Switch Image File
	ImageFileNames []string `json:"imageFileNames,omitempty"`

	// Name
	// Name of the Switch Model
	Name *string `json:"name,omitempty"`
}

func NewSwitchModel() *SwitchModel {
	switchModelType := new(SwitchModel)
	return switchModelType
}

func NewSwitchModelWithDefaults() *SwitchModel {
	switchModelType := new(SwitchModel)
	return switchModelType
}
