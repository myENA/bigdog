package vsz

// API Version: v8_1

import (
	"context"
	"encoding/json"
	"net/http"
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

func NewSwitchMFirmwaresQueryResultList() *SwitchMFirmwaresQueryResultList {
	m := new(SwitchMFirmwaresQueryResultList)
	return m
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

func NewSwitchMFirmwaresQueryResultListExtraType() *SwitchMFirmwaresQueryResultListExtraType {
	m := new(SwitchMFirmwaresQueryResultListExtraType)
	return m
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

func NewSwitchMFirmwareScheduleIds() *SwitchMFirmwareScheduleIds {
	m := new(SwitchMFirmwareScheduleIds)
	return m
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

func NewSwitchMFirmwareScheduleIdsExtraType() *SwitchMFirmwareScheduleIdsExtraType {
	m := new(SwitchMFirmwareScheduleIdsExtraType)
	return m
}

type SwitchMFirmwareSwitchFirmware struct {
	SwitchModels []*SwitchMFirmwareSwitchModel `json:"switchModels,omitempty"`

	// Version
	// Firmware version of the Switch
	Version *string `json:"version,omitempty"`
}

func NewSwitchMFirmwareSwitchFirmware() *SwitchMFirmwareSwitchFirmware {
	m := new(SwitchMFirmwareSwitchFirmware)
	return m
}

type SwitchMFirmwareSwitchModel struct {
	// ImageFileNames
	// Name of the Switch Image File
	ImageFileNames []string `json:"imageFileNames,omitempty"`

	// Name
	// Name of the Switch Model
	Name *string `json:"name,omitempty"`
}

func NewSwitchMFirmwareSwitchModel() *SwitchMFirmwareSwitchModel {
	m := new(SwitchMFirmwareSwitchModel)
	return m
}

// AddFirmware
//
// Use this API command to retrieve list of switch firmwares uploaded to SmartZone.
//
// Request Body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
func (s *SwitchMFirmwareService) AddFirmware(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet) (*SwitchMFirmwaresQueryResultList, error) {
	var (
		req  *APIRequest
		resp *SwitchMFirmwaresQueryResultList
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSwitchMAddFirmware, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
}

// AddFirmwareUpload
//
// Use this API command to upload a firmware image zip file to SmartZone.
//
// Request Body:
//	 - body []byte
func (s *SwitchMFirmwareService) AddFirmwareUpload(ctx context.Context, body []byte) (interface{}, error) {
	var (
		req  *APIRequest
		resp interface{}
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSwitchMAddFirmwareUpload, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
}

// DeleteFirmwareByVersion
//
// Use this API command to deletes a firmware image file from SmartZone.
//
// Required Parameters:
// - version string
//		- required
func (s *SwitchMFirmwareService) DeleteFirmwareByVersion(ctx context.Context, version string) (interface{}, error) {
	var (
		req  *APIRequest
		resp interface{}
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, version, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteSwitchMDeleteFirmwareByVersion, true)
	req.SetPathParameter("version", version)
}

// FindFirmware
//
// Use this API command to retrieve list of switch firmwares uploaded to SmartZone.
func (s *SwitchMFirmwareService) FindFirmware(ctx context.Context) (*SwitchMFirmwareAllFirmwaresQueryResultList, error) {
	var (
		req  *APIRequest
		resp *SwitchMFirmwareAllFirmwaresQueryResultList
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteSwitchMFindFirmware, true)
}

// PartialUpdateFirmwareByVersion
//
// Use this API command to update the given firmware version on switches matching criteria.
//
// Request Body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
//
// Required Parameters:
// - version string
//		- required
func (s *SwitchMFirmwareService) PartialUpdateFirmwareByVersion(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet, version string) (*SwitchMFirmwareScheduleIds, error) {
	var (
		req  *APIRequest
		resp *SwitchMFirmwareScheduleIds
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, version, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPatch, RouteSwitchMPartialUpdateFirmwareByVersion, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	req.SetPathParameter("version", version)
}
