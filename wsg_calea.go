package bigdog

// API Version: v9_1

import (
	"context"
	"errors"
	"io"
	"net/http"
	"time"
)

type WSGCALEAService struct {
	apiClient *VSZClient
}

func NewWSGCALEAService(c *VSZClient) *WSGCALEAService {
	s := new(WSGCALEAService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGCALEAService() *WSGCALEAService {
	return NewWSGCALEAService(ss.apiClient)
}

// WSGCALEACommonSettingRq
//
// Definition: calea_caleaCommonSettingRq
type WSGCALEACommonSettingRq struct {
	// CaleaServerIp
	// CARLEA Server IP
	CaleaServerIp *string `json:"caleaServerIp,omitempty"`

	// Dcip
	// DP IP in Data Center
	Dcip *string `json:"dc_ip,omitempty"`
}

func NewWSGCALEACommonSettingRq() *WSGCALEACommonSettingRq {
	m := new(WSGCALEACommonSettingRq)
	return m
}

// WSGCALEACommonSettingRsp
//
// Definition: calea_caleaCommonSettingRsp
type WSGCALEACommonSettingRsp struct {
	// CaleaServerIp
	// CARLEA Server IP
	CaleaServerIp *string `json:"caleaServerIp,omitempty"`

	// Dcip
	// DP IP in Data Center
	Dcip *string `json:"dc_ip,omitempty"`
}

type WSGCALEACommonSettingRspAPIResponse struct {
	*RawAPIResponse
	Data *WSGCALEACommonSettingRsp
}

func newWSGCALEACommonSettingRspAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGCALEACommonSettingRspAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGCALEACommonSettingRspAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGCALEACommonSettingRsp)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGCALEACommonSettingRsp() *WSGCALEACommonSettingRsp {
	m := new(WSGCALEACommonSettingRsp)
	return m
}

// WSGCALEAMacListRq
//
// Definition: calea_caleaMacListRq
type WSGCALEAMacListRq struct {
	MacList []string `json:"macList,omitempty"`
}

func NewWSGCALEAMacListRq() *WSGCALEAMacListRq {
	m := new(WSGCALEAMacListRq)
	return m
}

// WSGCALEAMacListRsp
//
// Definition: calea_caleaMacListRsp
type WSGCALEAMacListRsp struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []string `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGCALEAMacListRspAPIResponse struct {
	*RawAPIResponse
	Data *WSGCALEAMacListRsp
}

func newWSGCALEAMacListRspAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGCALEAMacListRspAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGCALEAMacListRspAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGCALEAMacListRsp)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGCALEAMacListRsp() *WSGCALEAMacListRsp {
	m := new(WSGCALEAMacListRsp)
	return m
}

// AddSystemCaleaCommonSetting
//
// Use this API command to set CALEA common setting.
//
// Operation ID: addSystemCaleaCommonSetting
// Operation path: /system/caleaCommonSetting
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGCALEACommonSettingRq
func (s *WSGCALEAService) AddSystemCaleaCommonSetting(ctx context.Context, body *WSGCALEACommonSettingRq, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGAddSystemCaleaCommonSetting, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*EmptyAPIResponse), err
}

// AddSystemCaleaMac
//
// Use this API command to add specified CALEA UE MACs
//
// Operation ID: addSystemCaleaMac
// Operation path: /system/caleaMac
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGCALEAMacListRq
func (s *WSGCALEAService) AddSystemCaleaMac(ctx context.Context, body *WSGCALEAMacListRq, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGAddSystemCaleaMac, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*EmptyAPIResponse), err
}

// AddSystemCaleaMacList
//
// Use this API command to upload csv file of CALEA UE MACs.
//
// Operation ID: addSystemCaleaMacList
// Operation path: /system/caleaMacList
// Success code: 200 (OK)
//
// Form data parameters:
// - uploadFile io.Reader
//		- required
func (s *WSGCALEAService) AddSystemCaleaMacList(ctx context.Context, filename string, uploadFile io.Reader, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGAddSystemCaleaMacList, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueMultipartFormData)
	req.Header.Set(headerKeyAccept, "*/*")
	req.MultipartForm()
	if err = req.AddMultipartFile("uploadFile", filename, uploadFile); err != nil {
		return resp.(*RawAPIResponse), err
	}
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*RawAPIResponse), err
}

// DeleteSystemCaleaMac
//
// Use this API command to delete specified CALEA UE MACs.
//
// Operation ID: deleteSystemCaleaMac
// Operation path: /system/caleaMac
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGCALEAMacListRq
func (s *WSGCALEAService) DeleteSystemCaleaMac(ctx context.Context, body *WSGCALEAMacListRq, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteSystemCaleaMac, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*EmptyAPIResponse), err
}

// DeleteSystemCaleaMacList
//
// Use this API command to delete all CALEA UE MACs.
//
// Operation ID: deleteSystemCaleaMacList
// Operation path: /system/caleaMacList
// Success code: 200 (OK)
func (s *WSGCALEAService) DeleteSystemCaleaMacList(ctx context.Context, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteSystemCaleaMacList, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*EmptyAPIResponse), err
}

// FindSystemCaleaCommonSetting
//
// Use this API command to get CALEA common setting.
//
// Operation ID: findSystemCaleaCommonSetting
// Operation path: /system/caleaCommonSetting
// Success code: 200 (OK)
func (s *WSGCALEAService) FindSystemCaleaCommonSetting(ctx context.Context, mutators ...RequestMutator) (*WSGCALEACommonSettingRspAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGCALEACommonSettingRspAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindSystemCaleaCommonSetting, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*WSGCALEACommonSettingRspAPIResponse), err
}

// FindSystemCaleaMacList
//
// Use this API command to get all CALEA UE MACs.
//
// Operation ID: findSystemCaleaMacList
// Operation path: /system/caleaMacList
// Success code: 200 (OK)
func (s *WSGCALEAService) FindSystemCaleaMacList(ctx context.Context, mutators ...RequestMutator) (*WSGCALEAMacListRspAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGCALEAMacListRspAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindSystemCaleaMacList, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*WSGCALEAMacListRspAPIResponse), err
}
