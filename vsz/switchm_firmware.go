package vsz

// API Version: v8_1

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"gopkg.in/go-playground/validator.v9"
)

type SwitchMFirmwareService struct {
	apiClient *APIClient
}

func NewSwitchMFirmwareService(c *APIClient) *SwitchMFirmwareService {
	s := new(SwitchMFirmwareService)
	s.apiClient = c
	return s
}

func (ss *SwitchMService) SwitchMFirmwareService() *SwitchMFirmwareService {
	return NewSwitchMFirmwareService(ss.apiClient)
}

type SwitchMFirmwareAllFirmwaresQueryResultList struct {
	*SwitchMFirmwaresQueryResultList
}

type SwitchMFirmwaresQueryResultList struct {
	// Extra
	// Extra information for Firmware list
	Extra *SwitchMFirmwaresQueryResultListExtraType `json:"extra,omitempty"`

	// FirstIndex
	// Index of the first firmware list returned out of the complete Firmware list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicator of whether there are more Firmwares after the current displayed list
	HasMore *bool `json:"hasMore,omitempty"`

	List []*SwitchMFirmwareSwitchFirmware `json:"list,omitempty"`

	// RawDataTotalCount
	// Firmware list count
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Total Firmware list count
	TotalCount *int `json:"totalCount,omitempty"`
}

// SwitchMFirmwaresQueryResultListExtraType
//
// Extra information for Firmware list
type SwitchMFirmwaresQueryResultListExtraType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMFirmwaresQueryResultListExtraType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMFirmwaresQueryResultListExtraType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMFirmwaresQueryResultListExtraType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

type SwitchMFirmwareScheduleIds struct {
	// Extra
	// Extra information for Schedule Ids list
	Extra *SwitchMFirmwareScheduleIdsExtraType `json:"extra,omitempty"`

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

// SwitchMFirmwareScheduleIdsExtraType
//
// Extra information for Schedule Ids list
type SwitchMFirmwareScheduleIdsExtraType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMFirmwareScheduleIdsExtraType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMFirmwareScheduleIdsExtraType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMFirmwareScheduleIdsExtraType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

type SwitchMFirmwareSwitchFirmware struct {
	SwitchModels []*SwitchMFirmwareSwitchModel `json:"switchModels,omitempty"`

	// Version
	// Firmware version of the Switch
	Version *string `json:"version,omitempty"`
}

type SwitchMFirmwareSwitchModel struct {
	// ImageFileNames
	// Name of the Switch Image File
	ImageFileNames []string `json:"imageFileNames,omitempty"`

	// Name
	// Name of the Switch Model
	Name *string `json:"name,omitempty"`
}

// AddFirmware
//
// Use this API command to retrieve list of switch firmwares uploaded to SmartZone.
//
// Request Body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
func (s *SwitchMFirmwareService) AddFirmware(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet) (*SwitchMFirmwaresQueryResultList, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// AddFirmwareUpload
//
// Use this API command to upload a firmware image zip file to SmartZone.
func (s *SwitchMFirmwareService) AddFirmwareUpload(ctx context.Context) error {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteFirmwareByVersion
//
// Use this API command to deletes a firmware image file from SmartZone.
//
// Path Parameters:
// - pVersion string
//		- required
func (s *SwitchMFirmwareService) DeleteFirmwareByVersion(ctx context.Context, pVersion string) error {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
}

// FindFirmware
//
// Use this API command to retrieve list of switch firmwares uploaded to SmartZone.
func (s *SwitchMFirmwareService) FindFirmware(ctx context.Context) (*SwitchMFirmwareAllFirmwaresQueryResultList, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// PartialUpdateFirmwareByVersion
//
// Use this API command to update the given firmware version on switches matching criteria.
//
// Request Body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
//
// Path Parameters:
// - pVersion string
//		- required
func (s *SwitchMFirmwareService) PartialUpdateFirmwareByVersion(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet, pVersion string) (*SwitchMFirmwareScheduleIds, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}
