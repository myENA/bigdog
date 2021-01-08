package bigdog

// API Version: v9_1

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

// SwitchMFirmwareConfigFirmwaresQueryResultList
//
// Definition: firmware_firmwaresQueryResultList
type SwitchMFirmwareConfigFirmwaresQueryResultList struct {
	// Extra
	// Extra information for Firmware list
	Extra interface{} `json:"extra,omitempty"`

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

type SwitchMFirmwareConfigFirmwaresQueryResultListAPIResponse struct {
	*RawAPIResponse
	Data *SwitchMFirmwareConfigFirmwaresQueryResultList
}

func newSwitchMFirmwareConfigFirmwaresQueryResultListAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SwitchMFirmwareConfigFirmwaresQueryResultListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SwitchMFirmwareConfigFirmwaresQueryResultListAPIResponse) Hydrate() error {
	r.Data = new(SwitchMFirmwareConfigFirmwaresQueryResultList)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSwitchMFirmwareConfigFirmwaresQueryResultList() *SwitchMFirmwareConfigFirmwaresQueryResultList {
	m := new(SwitchMFirmwareConfigFirmwaresQueryResultList)
	return m
}

// SwitchMFirmwareConfigScheduleIds
//
// Definition: firmware_scheduleIds
type SwitchMFirmwareConfigScheduleIds struct {
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

type SwitchMFirmwareConfigScheduleIdsAPIResponse struct {
	*RawAPIResponse
	Data *SwitchMFirmwareConfigScheduleIds
}

func newSwitchMFirmwareConfigScheduleIdsAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SwitchMFirmwareConfigScheduleIdsAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SwitchMFirmwareConfigScheduleIdsAPIResponse) Hydrate() error {
	r.Data = new(SwitchMFirmwareConfigScheduleIds)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSwitchMFirmwareConfigScheduleIds() *SwitchMFirmwareConfigScheduleIds {
	m := new(SwitchMFirmwareConfigScheduleIds)
	return m
}

// SwitchMFirmwareConfigSwitchFirmware
//
// Definition: firmware_switchFirmware
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

// SwitchMFirmwareConfigSwitchModel
//
// Definition: firmware_switchModel
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
// Operation ID: addFirmware
//
// Use this API command to retrieve list of switch firmwares uploaded to SmartZone.
//
// Request Body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
func (s *SwitchMFirmwareConfigService) AddFirmware(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*SwitchMFirmwareConfigFirmwaresQueryResultListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSwitchMFirmwareConfigFirmwaresQueryResultListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SwitchMFirmwareConfigFirmwaresQueryResultListAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteSwitchMAddFirmware, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*SwitchMFirmwareConfigFirmwaresQueryResultListAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SwitchMFirmwareConfigFirmwaresQueryResultListAPIResponse), err
}

// AddFirmwareUpload
//
// Operation ID: addFirmwareUpload
//
// Use this API command to upload a firmware image zip file to SmartZone.
//
// Form Data Parameters:
// - uploadFile io.Reader
//		- required
func (s *SwitchMFirmwareConfigService) AddFirmwareUpload(ctx context.Context, filename string, uploadFile io.Reader, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteSwitchMAddFirmwareUpload, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueMultipartFormData)
	req.Header.Set(headerKeyAccept, "*/*")
	req.MultipartForm()
	if err = req.AddMultipartFile("uploadFile", filename, uploadFile); err != nil {
		return resp.(*RawAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// DeleteFirmwareByVersion
//
// Operation ID: deleteFirmwareByVersion
//
// Use this API command to deletes a firmware image file from SmartZone.
//
// Required Parameters:
// - version string
//		- required
func (s *SwitchMFirmwareConfigService) DeleteFirmwareByVersion(ctx context.Context, version string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodDelete, RouteSwitchMDeleteFirmwareByVersion, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("version", version)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// FindFirmware
//
// Operation ID: findFirmware
//
// Use this API command to retrieve list of switch firmwares uploaded to SmartZone.
func (s *SwitchMFirmwareConfigService) FindFirmware(ctx context.Context, mutators ...RequestMutator) (*SwitchMFirmwareConfigFirmwaresQueryResultListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSwitchMFirmwareConfigFirmwaresQueryResultListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SwitchMFirmwareConfigFirmwaresQueryResultListAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteSwitchMFindFirmware, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SwitchMFirmwareConfigFirmwaresQueryResultListAPIResponse), err
}

// PartialUpdateFirmwareByVersion
//
// Operation ID: partialUpdateFirmwareByVersion
//
// Use this API command to update the given firmware version on switches matching criteria.
//
// Request Body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
//
// Required Parameters:
// - version string
//		- required
func (s *SwitchMFirmwareConfigService) PartialUpdateFirmwareByVersion(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet, version string, mutators ...RequestMutator) (*SwitchMFirmwareConfigScheduleIdsAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSwitchMFirmwareConfigScheduleIdsAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SwitchMFirmwareConfigScheduleIdsAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPatch, RouteSwitchMPartialUpdateFirmwareByVersion, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*SwitchMFirmwareConfigScheduleIdsAPIResponse), err
	}
	req.PathParams.Set("version", version)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SwitchMFirmwareConfigScheduleIdsAPIResponse), err
}
