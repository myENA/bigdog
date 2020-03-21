package vsz

// API Version: v9_0

import (
	"context"
	"net/http"
)

type WSGDPNATProfileService struct {
	apiClient *APIClient
}

func NewWSGDPNATProfileService(c *APIClient) *WSGDPNATProfileService {
	s := new(WSGDPNATProfileService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGDPNATProfileService() *WSGDPNATProfileService {
	return NewWSGDPNATProfileService(ss.apiClient)
}

// AddDpNatProfiles
//
// Use this API command to create DHCP NAT profile - basic.
//
// Request Body:
//	 - body *WSGDPProfileDpNatProfileBasicBO
func (s *WSGDPNATProfileService) AddDpNatProfiles(ctx context.Context, body *WSGDPProfileDpNatProfileBasicBO) (*WSGDPProfileDpNatProfileBasicBO, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGDPProfileDpNatProfileBasicBO
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
	req = NewAPIRequest(http.MethodPost, RouteWSGAddDpNatProfiles, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGDPProfileDpNatProfileBasicBO()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// AddDpNatProfilesDpNatPoolsById
//
// Use this API command to create DHCP NAT profile - pool.
//
// Request Body:
//	 - body *WSGDPProfileDpNatProfilePoolBO
//
// Required Parameters:
// - id string
//		- required
func (s *WSGDPNATProfileService) AddDpNatProfilesDpNatPoolsById(ctx context.Context, body *WSGDPProfileDpNatProfilePoolBO, id string) (*WSGDPProfileDpNatProfilePoolBO, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGDPProfileDpNatProfilePoolBO
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
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddDpNatProfilesDpNatPoolsById, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGDPProfileDpNatProfilePoolBO()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// DeleteDpNatProfiles
//
// Use this API command to delete DHCP NAT profiles.
//
// Request Body:
//	 - body *WSGDPProfileBulkDelete
func (s *WSGDPNATProfileService) DeleteDpNatProfiles(ctx context.Context, body *WSGDPProfileBulkDelete) (interface{}, *APIResponseMeta, error) {
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
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteDpNatProfiles, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = new(interface{})
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// DeleteDpNatProfilesById
//
// Use this API command to delete DHCP NAT profile.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGDPNATProfileService) DeleteDpNatProfilesById(ctx context.Context, id string) (interface{}, *APIResponseMeta, error) {
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
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteDpNatProfilesById, true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = new(interface{})
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// DeleteDpNatProfilesDpNatPoolsById
//
// Use this API command to delete DP NAT profile - pools.
//
// Request Body:
//	 - body *WSGDPProfileBulkDelete
//
// Required Parameters:
// - id string
//		- required
func (s *WSGDPNATProfileService) DeleteDpNatProfilesDpNatPoolsById(ctx context.Context, body *WSGDPProfileBulkDelete, id string) (interface{}, *APIResponseMeta, error) {
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
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteDpNatProfilesDpNatPoolsById, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = new(interface{})
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// DeleteDpNatProfilesDpNatPoolsByPoolId
//
// Use this API command to delete DP NAT profile - pool.
//
// Required Parameters:
// - id string
//		- required
// - poolId string
//		- required
func (s *WSGDPNATProfileService) DeleteDpNatProfilesDpNatPoolsByPoolId(ctx context.Context, id string, poolId string) (interface{}, *APIResponseMeta, error) {
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
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, poolId, "required"); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteDpNatProfilesDpNatPoolsByPoolId, true)
	req.SetPathParameter("id", id)
	req.SetPathParameter("poolId", poolId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = new(interface{})
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindDpNatProfiles
//
// Use this API command to retrieve DHCP NAT profile - basic list.
func (s *WSGDPNATProfileService) FindDpNatProfiles(ctx context.Context) (*WSGDPProfileDpNatProfileBasicBOList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGDPProfileDpNatProfileBasicBOList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindDpNatProfiles, true)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGDPProfileDpNatProfileBasicBOList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// FindDpNatProfilesById
//
// Use this API command to retrieve DHCP NAT profile - basic.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGDPNATProfileService) FindDpNatProfilesById(ctx context.Context, id string) (*WSGDPProfileDpNatProfileBasicBO, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGDPProfileDpNatProfileBasicBO
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindDpNatProfilesById, true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGDPProfileDpNatProfileBasicBO()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// FindDpNatProfilesDpNatPoolsById
//
// Use this API command to retrieve DP NAT profile - pool list.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGDPNATProfileService) FindDpNatProfilesDpNatPoolsById(ctx context.Context, id string) (*WSGDPProfileDpNatProfilePoolBOList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGDPProfileDpNatProfilePoolBOList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindDpNatProfilesDpNatPoolsById, true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGDPProfileDpNatProfilePoolBOList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// FindDpNatProfilesDpNatPoolsByPoolId
//
// Use this API command to retrieve DP DHCP profile - pool.
//
// Required Parameters:
// - id string
//		- required
// - poolId string
//		- required
func (s *WSGDPNATProfileService) FindDpNatProfilesDpNatPoolsByPoolId(ctx context.Context, id string, poolId string) (*WSGDPProfileDpNatProfilePoolBO, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGDPProfileDpNatProfilePoolBO
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, poolId, "required"); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindDpNatProfilesDpNatPoolsByPoolId, true)
	req.SetPathParameter("id", id)
	req.SetPathParameter("poolId", poolId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGDPProfileDpNatProfilePoolBO()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// UpdateDpNatProfilesById
//
// Use this API command to modify DHCP NAT profile - basic.
//
// Request Body:
//	 - body *WSGDPProfileDpNatProfileBasicBO
//
// Required Parameters:
// - id string
//		- required
func (s *WSGDPNATProfileService) UpdateDpNatProfilesById(ctx context.Context, body *WSGDPProfileDpNatProfileBasicBO, id string) (*WSGDPProfileDpNatProfileBasicBO, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGDPProfileDpNatProfileBasicBO
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
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPut, RouteWSGUpdateDpNatProfilesById, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGDPProfileDpNatProfileBasicBO()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// UpdateDpNatProfilesDpNatPoolsByPoolId
//
// Use this API command to modify DHCP NAT profile - pool.
//
// Request Body:
//	 - body *WSGDPProfileDpNatProfilePoolBO
//
// Required Parameters:
// - id string
//		- required
// - poolId string
//		- required
func (s *WSGDPNATProfileService) UpdateDpNatProfilesDpNatPoolsByPoolId(ctx context.Context, body *WSGDPProfileDpNatProfilePoolBO, id string, poolId string) (*WSGDPProfileDpNatProfilePoolBO, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGDPProfileDpNatProfilePoolBO
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
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, poolId, "required"); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPut, RouteWSGUpdateDpNatProfilesDpNatPoolsByPoolId, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetPathParameter("id", id)
	req.SetPathParameter("poolId", poolId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGDPProfileDpNatProfilePoolBO()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}
