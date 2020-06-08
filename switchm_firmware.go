package ruckus

// API Version: v9_0

import (
	"context"
	"encoding/json"
	"net/http"
)

type SwitchMSwitchFirmwareConfigService struct {
	apiClient *VSZClient
}

func NewSwitchMSwitchFirmwareConfigService(c *VSZClient) *SwitchMSwitchFirmwareConfigService {
	s := new(SwitchMSwitchFirmwareConfigService)
	s.apiClient = c
	return s
}

func (ss *SwitchMService) SwitchMSwitchFirmwareConfigService() *SwitchMSwitchFirmwareConfigService {
	return NewSwitchMSwitchFirmwareConfigService(ss.apiClient)
}

type SwitchMSwitchFirmwareConfigFirmwaresQueryResultList struct {
	// Extra
	// Extra information for Firmware list
	Extra *SwitchMSwitchFirmwareConfigFirmwaresQueryResultList `json:"extra,omitempty"`

	// FirstIndex
	// Index of the first firmware list returned out of the complete Firmware list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicator of whether there are more Firmwares after the current displayed list
	HasMore *bool `json:"hasMore,omitempty"`

	List []*SwitchMSwitchFirmwareConfigFirmwaresQueryResultList `json:"list,omitempty"`

	// RawDataTotalCount
	// Firmware list count
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Total Firmware list count
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewSwitchMSwitchFirmwareConfigFirmwaresQueryResultList() *SwitchMSwitchFirmwareConfigFirmwaresQueryResultList {
	m := new(SwitchMSwitchFirmwareConfigFirmwaresQueryResultList)
	return m
}

// SwitchMSwitchFirmwareConfigFirmwaresQueryResultListExtraType
//
// Extra information for Firmware list
type SwitchMSwitchFirmwareConfigFirmwaresQueryResultListExtraType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMSwitchFirmwareConfigFirmwaresQueryResultListExtraType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMSwitchFirmwareConfigFirmwaresQueryResultListExtraType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMSwitchFirmwareConfigFirmwaresQueryResultListExtraType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSwitchMSwitchFirmwareConfigFirmwaresQueryResultListExtraType() *SwitchMSwitchFirmwareConfigFirmwaresQueryResultListExtraType {
	m := new(SwitchMSwitchFirmwareConfigFirmwaresQueryResultListExtraType)
	return m
}

type SwitchMSwitchFirmwareConfigScheduleIds struct {
	// Extra
	// Extra information for Schedule Ids list
	Extra *SwitchMSwitchFirmwareConfigScheduleIds `json:"extra,omitempty"`

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

func NewSwitchMSwitchFirmwareConfigScheduleIds() *SwitchMSwitchFirmwareConfigScheduleIds {
	m := new(SwitchMSwitchFirmwareConfigScheduleIds)
	return m
}

// SwitchMSwitchFirmwareConfigScheduleIdsExtraType
//
// Extra information for Schedule Ids list
type SwitchMSwitchFirmwareConfigScheduleIdsExtraType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMSwitchFirmwareConfigScheduleIdsExtraType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMSwitchFirmwareConfigScheduleIdsExtraType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMSwitchFirmwareConfigScheduleIdsExtraType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSwitchMSwitchFirmwareConfigScheduleIdsExtraType() *SwitchMSwitchFirmwareConfigScheduleIdsExtraType {
	m := new(SwitchMSwitchFirmwareConfigScheduleIdsExtraType)
	return m
}

type SwitchMSwitchFirmwareConfigSwitchFirmware struct {
	SwitchModels []*SwitchMSwitchFirmwareConfigSwitchFirmware `json:"switchModels,omitempty"`

	// Version
	// Firmware version of the Switch
	Version *string `json:"version,omitempty"`
}

func NewSwitchMSwitchFirmwareConfigSwitchFirmware() *SwitchMSwitchFirmwareConfigSwitchFirmware {
	m := new(SwitchMSwitchFirmwareConfigSwitchFirmware)
	return m
}

type SwitchMSwitchFirmwareConfigSwitchModel struct {
	// ImageFileNames
	// Name of the Switch Image File
	ImageFileNames []string `json:"imageFileNames,omitempty"`

	// Name
	// Name of the Switch Model
	Name *string `json:"name,omitempty"`
}

func NewSwitchMSwitchFirmwareConfigSwitchModel() *SwitchMSwitchFirmwareConfigSwitchModel {
	m := new(SwitchMSwitchFirmwareConfigSwitchModel)
	return m
}

// AddFirmware
//
// Use this API command to retrieve list of switch firmwares uploaded to SmartZone.
//
// Request Body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
func (s *SwitchMSwitchFirmwareConfigService) AddFirmware(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet) (*SwitchMSwitchFirmwareConfigFirmwaresQueryResultList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMSwitchFirmwareConfigFirmwaresQueryResultList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSwitchMAddFirmware, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMSwitchFirmwareConfigFirmwaresQueryResultList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// AddFirmwareUpload
//
// Use this API command to upload a firmware image zip file to SmartZone.
//
// Request Body:
//	 - body []byte
func (s *SwitchMSwitchFirmwareConfigService) AddFirmwareUpload(ctx context.Context, body []byte) (interface{}, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     interface{}
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSwitchMAddFirmwareUpload, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = new(interface{})
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// DeleteFirmwareByVersion
//
// Use this API command to deletes a firmware image file from SmartZone.
//
// Required Parameters:
// - version string
//		- required
func (s *SwitchMSwitchFirmwareConfigService) DeleteFirmwareByVersion(ctx context.Context, version string) (interface{}, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     interface{}
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteSwitchMDeleteFirmwareByVersion, true)
	req.SetPathParameter("version", version)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = new(interface{})
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindFirmware
//
// Use this API command to retrieve list of switch firmwares uploaded to SmartZone.
func (s *SwitchMSwitchFirmwareConfigService) FindFirmware(ctx context.Context) (*SwitchMSwitchFirmwareConfigFirmwaresQueryResultList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMSwitchFirmwareConfigFirmwaresQueryResultList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteSwitchMFindFirmware, true)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMSwitchFirmwareConfigFirmwaresQueryResultList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
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
func (s *SwitchMSwitchFirmwareConfigService) PartialUpdateFirmwareByVersion(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet, version string) (*SwitchMSwitchFirmwareConfigScheduleIds, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMSwitchFirmwareConfigScheduleIds
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPatch, RouteSwitchMPartialUpdateFirmwareByVersion, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetPathParameter("version", version)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMSwitchFirmwareConfigScheduleIds()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
