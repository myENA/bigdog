package bigdog

// API Version: v9_1

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
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

func newWSGCALEACommonSettingRspAPIResponse(req *APIRequest, successCode int, httpResp *http.Response) APIResponse {
	r := new(WSGCALEACommonSettingRspAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(req, successCode, httpResp).(*RawAPIResponse)
	return r
}

func (r *WSGCALEACommonSettingRspAPIResponse) Hydrate() error {
	r.Data = new(WSGCALEACommonSettingRsp)
	return json.NewDecoder(r).Decode(r.Data)
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

func newWSGCALEAMacListRspAPIResponse(req *APIRequest, successCode int, httpResp *http.Response) APIResponse {
	r := new(WSGCALEAMacListRspAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(req, successCode, httpResp).(*RawAPIResponse)
	return r
}

func (r *WSGCALEAMacListRspAPIResponse) Hydrate() error {
	r.Data = new(WSGCALEAMacListRsp)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewWSGCALEAMacListRsp() *WSGCALEAMacListRsp {
	m := new(WSGCALEAMacListRsp)
	return m
}

// AddSystemCaleaCommonSetting
//
// Operation ID: addSystemCaleaCommonSetting
//
// Use this API command to set CALEA common setting.
//
// Request Body:
//	 - body *WSGCALEACommonSettingRq
func (s *WSGCALEAService) AddSystemCaleaCommonSetting(ctx context.Context, body *WSGCALEACommonSettingRq, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodPost, RouteWSGAddSystemCaleaCommonSetting, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*RawAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// AddSystemCaleaMac
//
// Operation ID: addSystemCaleaMac
//
// Use this API command to add specified CALEA UE MACs
//
// Request Body:
//	 - body *WSGCALEAMacListRq
func (s *WSGCALEAService) AddSystemCaleaMac(ctx context.Context, body *WSGCALEAMacListRq, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodPost, RouteWSGAddSystemCaleaMac, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*RawAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// AddSystemCaleaMacList
//
// Operation ID: addSystemCaleaMacList
//
// Use this API command to upload csv file of CALEA UE MACs.
//
// Form Data Parameters:
// - uploadFile io.Reader
//		- required
func (s *WSGCALEAService) AddSystemCaleaMacList(ctx context.Context, filename string, uploadFile io.Reader, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodPost, RouteWSGAddSystemCaleaMacList, true)
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

// DeleteSystemCaleaMac
//
// Operation ID: deleteSystemCaleaMac
//
// Use this API command to delete specified CALEA UE MACs.
//
// Request Body:
//	 - body *WSGCALEAMacListRq
func (s *WSGCALEAService) DeleteSystemCaleaMac(ctx context.Context, body *WSGCALEAMacListRq, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteSystemCaleaMac, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*RawAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// DeleteSystemCaleaMacList
//
// Operation ID: deleteSystemCaleaMacList
//
// Use this API command to delete all CALEA UE MACs.
func (s *WSGCALEAService) DeleteSystemCaleaMacList(ctx context.Context, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteSystemCaleaMacList, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// FindSystemCaleaCommonSetting
//
// Operation ID: findSystemCaleaCommonSetting
//
// Use this API command to get CALEA common setting.
func (s *WSGCALEAService) FindSystemCaleaCommonSetting(ctx context.Context, mutators ...RequestMutator) (*WSGCALEACommonSettingRspAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGCALEACommonSettingRspAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGCALEACommonSettingRspAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindSystemCaleaCommonSetting, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGCALEACommonSettingRspAPIResponse), err
}

// FindSystemCaleaMacList
//
// Operation ID: findSystemCaleaMacList
//
// Use this API command to get all CALEA UE MACs.
func (s *WSGCALEAService) FindSystemCaleaMacList(ctx context.Context, mutators ...RequestMutator) (*WSGCALEAMacListRspAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGCALEAMacListRspAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGCALEAMacListRspAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindSystemCaleaMacList, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGCALEAMacListRspAPIResponse), err
}
