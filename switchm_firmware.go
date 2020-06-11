package bigdog

// API Version: v9_0

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
)

type SwitchMFirmwareConfigService struct {
	apiClient *VSZClient
}

func NewSwitchMFirmwareConfigService(c *VSZClient) *SwitchMFirmwareConfigService {
	s := new(SwitchMFirmwareConfigService)
	s.apiClient = c
	return s
}

func (ss *SwitchMService) SwitchMFirmwareConfigService() *SwitchMFirmwareConfigService {
	return NewSwitchMFirmwareConfigService(ss.apiClient)
}

type SwitchMFirmwareConfigFirmwaresQueryResultList struct {
	// Extra
	// Extra information for Firmware list
	Extra *SwitchMFirmwareConfigFirmwaresQueryResultListExtraType `json:"extra,omitempty"`

	// FirstIndex
	// Index of the first firmware list returned out of the complete Firmware list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicator of whether there are more Firmwares after the current displayed list
	HasMore *bool `json:"hasMore,omitempty"`

	List []*SwitchMFirmwareConfigSwitchFirmware `json:"list,omitempty"`

	// RawDataTotalCount
	// Firmware list count
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Total Firmware list count
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewSwitchMFirmwareConfigFirmwaresQueryResultList() *SwitchMFirmwareConfigFirmwaresQueryResultList {
	m := new(SwitchMFirmwareConfigFirmwaresQueryResultList)
	return m
}

// SwitchMFirmwareConfigFirmwaresQueryResultListExtraType
//
// Extra information for Firmware list
type SwitchMFirmwareConfigFirmwaresQueryResultListExtraType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMFirmwareConfigFirmwaresQueryResultListExtraType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMFirmwareConfigFirmwaresQueryResultListExtraType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMFirmwareConfigFirmwaresQueryResultListExtraType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSwitchMFirmwareConfigFirmwaresQueryResultListExtraType() *SwitchMFirmwareConfigFirmwaresQueryResultListExtraType {
	m := new(SwitchMFirmwareConfigFirmwaresQueryResultListExtraType)
	return m
}

type SwitchMFirmwareConfigScheduleIds struct {
	// Extra
	// Extra information for Schedule Ids list
	Extra *SwitchMFirmwareConfigScheduleIdsExtraType `json:"extra,omitempty"`

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

func NewSwitchMFirmwareConfigScheduleIds() *SwitchMFirmwareConfigScheduleIds {
	m := new(SwitchMFirmwareConfigScheduleIds)
	return m
}

// SwitchMFirmwareConfigScheduleIdsExtraType
//
// Extra information for Schedule Ids list
type SwitchMFirmwareConfigScheduleIdsExtraType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMFirmwareConfigScheduleIdsExtraType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMFirmwareConfigScheduleIdsExtraType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMFirmwareConfigScheduleIdsExtraType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSwitchMFirmwareConfigScheduleIdsExtraType() *SwitchMFirmwareConfigScheduleIdsExtraType {
	m := new(SwitchMFirmwareConfigScheduleIdsExtraType)
	return m
}

type SwitchMFirmwareConfigSwitchFirmware struct {
	SwitchModels []*SwitchMFirmwareConfigSwitchModel `json:"switchModels,omitempty"`

	// Version
	// Firmware version of the Switch
	Version *string `json:"version,omitempty"`
}

func NewSwitchMFirmwareConfigSwitchFirmware() *SwitchMFirmwareConfigSwitchFirmware {
	m := new(SwitchMFirmwareConfigSwitchFirmware)
	return m
}

type SwitchMFirmwareConfigSwitchModel struct {
	// ImageFileNames
	// Name of the Switch Image File
	ImageFileNames []string `json:"imageFileNames,omitempty"`

	// Name
	// Name of the Switch Model
	Name *string `json:"name,omitempty"`
}

func NewSwitchMFirmwareConfigSwitchModel() *SwitchMFirmwareConfigSwitchModel {
	m := new(SwitchMFirmwareConfigSwitchModel)
	return m
}

// AddFirmware
//
// Use this API command to retrieve list of switch firmwares uploaded to SmartZone.
//
// Request Body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
func (s *SwitchMFirmwareConfigService) AddFirmware(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*SwitchMFirmwareConfigFirmwaresQueryResultList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMFirmwareConfigFirmwaresQueryResultList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSwitchMAddFirmware, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSwitchMFirmwareConfigFirmwaresQueryResultList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// AddFirmwareUpload
//
// Use this API command to upload a firmware image zip file to SmartZone.
//
// Form Data Parameters:
// - uploadFile io.Reader
//		- required
func (s *SwitchMFirmwareConfigService) AddFirmwareUpload(ctx context.Context, uploadFile io.Reader, mutators ...RequestMutator) (interface{}, *APIResponseMeta, error) {
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
	req.SetHeader(headerKeyContentType, headerValueMultipartFormData)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = AddRequestMultipartValues(req, map[string]interface{}{"uploadFile": uploadFile}); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
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
func (s *SwitchMFirmwareConfigService) DeleteFirmwareByVersion(ctx context.Context, version string, mutators ...RequestMutator) (interface{}, *APIResponseMeta, error) {
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
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	req.SetPathParameter("version", version)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = new(interface{})
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindFirmware
//
// Use this API command to retrieve list of switch firmwares uploaded to SmartZone.
func (s *SwitchMFirmwareConfigService) FindFirmware(ctx context.Context, mutators ...RequestMutator) (*SwitchMFirmwareConfigFirmwaresQueryResultList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMFirmwareConfigFirmwaresQueryResultList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteSwitchMFindFirmware, true)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSwitchMFirmwareConfigFirmwaresQueryResultList()
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
func (s *SwitchMFirmwareConfigService) PartialUpdateFirmwareByVersion(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet, version string, mutators ...RequestMutator) (*SwitchMFirmwareConfigScheduleIds, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMFirmwareConfigScheduleIds
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPatch, RouteSwitchMPartialUpdateFirmwareByVersion, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetPathParameter("version", version)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSwitchMFirmwareConfigScheduleIds()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
