package vsz

// API Version: v8_1

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
func (s *WSGCALEAService) AddSystemCaleaCommonSetting(ctx context.Context, body *WSGCALEACommonSettingRq) (*WSGCommonEmptyResult, error) {
	var (
		req  *APIRequest
		resp *WSGCommonEmptyResult
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
	req = NewAPIRequest(http.MethodPost, RouteWSGAddSystemCaleaCommonSetting, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
}

// AddSystemCaleaMac
//
// Use this API command to add specified CALEA UE MACs
//
// Request Body:
//	 - body *WSGCALEAMacListRq
func (s *WSGCALEAService) AddSystemCaleaMac(ctx context.Context, body *WSGCALEAMacListRq) (*WSGCommonEmptyResult, error) {
	var (
		req  *APIRequest
		resp *WSGCommonEmptyResult
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
	req = NewAPIRequest(http.MethodPost, RouteWSGAddSystemCaleaMac, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
}

// AddSystemCaleaMacList
//
// Use this API command to upload csv file of CALEA UE MACs.
//
// Request Body:
//	 - body []byte
func (s *WSGCALEAService) AddSystemCaleaMacList(ctx context.Context, body []byte) (interface{}, error) {
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
	req = NewAPIRequest(http.MethodPost, RouteWSGAddSystemCaleaMacList, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
}

// DeleteSystemCaleaMac
//
// Use this API command to delete specified CALEA UE MACs.
//
// Request Body:
//	 - body *WSGCALEAMacListRq
func (s *WSGCALEAService) DeleteSystemCaleaMac(ctx context.Context, body *WSGCALEAMacListRq) (*WSGCommonEmptyResult, error) {
	var (
		req  *APIRequest
		resp *WSGCommonEmptyResult
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
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteSystemCaleaMac, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
}

// DeleteSystemCaleaMacList
//
// Use this API command to delete all CALEA UE MACs.
func (s *WSGCALEAService) DeleteSystemCaleaMacList(ctx context.Context) (*WSGCommonEmptyResult, error) {
	var (
		req  *APIRequest
		resp *WSGCommonEmptyResult
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteSystemCaleaMacList, true)
}

// FindSystemCaleaCommonSetting
//
// Use this API command to get CALEA common setting.
func (s *WSGCALEAService) FindSystemCaleaCommonSetting(ctx context.Context) (*WSGCALEACommonSettingRsp, error) {
	var (
		req  *APIRequest
		resp *WSGCALEACommonSettingRsp
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindSystemCaleaCommonSetting, true)
}

// FindSystemCaleaMacList
//
// Use this API command to get all CALEA UE MACs.
func (s *WSGCALEAService) FindSystemCaleaMacList(ctx context.Context) (*WSGCALEAMacListRsp, error) {
	var (
		req  *APIRequest
		resp *WSGCALEAMacListRsp
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindSystemCaleaMacList, true)
}
