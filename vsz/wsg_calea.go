package vsz

// API Version: v9_0

import (
	"context"
	"net/http"
)

type WSGCALEAService struct {
	apiClient *APIClient
}

func NewWSGCALEAService(c *APIClient) *WSGCALEAService {
	s := new(WSGCALEAService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGCALEAService() *WSGCALEAService {
	return NewWSGCALEAService(ss.apiClient)
}

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

type WSGCALEACommonSettingRsp struct {
	// CaleaServerIp
	// CARLEA Server IP
	CaleaServerIp *string `json:"caleaServerIp,omitempty"`

	// Dcip
	// DP IP in Data Center
	Dcip *string `json:"dc_ip,omitempty"`
}

func NewWSGCALEACommonSettingRsp() *WSGCALEACommonSettingRsp {
	m := new(WSGCALEACommonSettingRsp)
	return m
}

type WSGCALEAMacListRq struct {
	MacList []string `json:"macList,omitempty"`
}

func NewWSGCALEAMacListRq() *WSGCALEAMacListRq {
	m := new(WSGCALEAMacListRq)
	return m
}

type WSGCALEAMacListRsp struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []string `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGCALEAMacListRsp() *WSGCALEAMacListRsp {
	m := new(WSGCALEAMacListRsp)
	return m
}

// AddSystemCaleaCommonSetting
//
// Use this API command to set CALEA common setting.
//
// Request Body:
//	 - body *WSGCALEACommonSettingRq
func (s *WSGCALEAService) AddSystemCaleaCommonSetting(ctx context.Context, body *WSGCALEACommonSettingRq) (*WSGCommonEmptyResult, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGCommonEmptyResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, rm, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddSystemCaleaCommonSetting, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonEmptyResult()
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, &resp, err)
	return resp, rm, err
}

// AddSystemCaleaMac
//
// Use this API command to add specified CALEA UE MACs
//
// Request Body:
//	 - body *WSGCALEAMacListRq
func (s *WSGCALEAService) AddSystemCaleaMac(ctx context.Context, body *WSGCALEAMacListRq) (*WSGCommonEmptyResult, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGCommonEmptyResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, rm, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddSystemCaleaMac, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonEmptyResult()
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, &resp, err)
	return resp, rm, err
}

// AddSystemCaleaMacList
//
// Use this API command to upload csv file of CALEA UE MACs.
//
// Request Body:
//	 - body []byte
func (s *WSGCALEAService) AddSystemCaleaMacList(ctx context.Context, body []byte) (interface{}, *APIResponseMeta, error) {
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
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddSystemCaleaMacList, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = new(interface{})
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// DeleteSystemCaleaMac
//
// Use this API command to delete specified CALEA UE MACs.
//
// Request Body:
//	 - body *WSGCALEAMacListRq
func (s *WSGCALEAService) DeleteSystemCaleaMac(ctx context.Context, body *WSGCALEAMacListRq) (*WSGCommonEmptyResult, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGCommonEmptyResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, rm, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteSystemCaleaMac, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonEmptyResult()
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, &resp, err)
	return resp, rm, err
}

// DeleteSystemCaleaMacList
//
// Use this API command to delete all CALEA UE MACs.
func (s *WSGCALEAService) DeleteSystemCaleaMacList(ctx context.Context) (*WSGCommonEmptyResult, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGCommonEmptyResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteSystemCaleaMacList, true)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonEmptyResult()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// FindSystemCaleaCommonSetting
//
// Use this API command to get CALEA common setting.
func (s *WSGCALEAService) FindSystemCaleaCommonSetting(ctx context.Context) (*WSGCALEACommonSettingRsp, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGCALEACommonSettingRsp
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindSystemCaleaCommonSetting, true)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCALEACommonSettingRsp()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// FindSystemCaleaMacList
//
// Use this API command to get all CALEA UE MACs.
func (s *WSGCALEAService) FindSystemCaleaMacList(ctx context.Context) (*WSGCALEAMacListRsp, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGCALEAMacListRsp
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindSystemCaleaMacList, true)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCALEAMacListRsp()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}
