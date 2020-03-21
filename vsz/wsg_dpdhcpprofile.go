package vsz

// API Version: v9_0

import (
	"context"
	"net/http"
)

type WSGDPDHCPProfileService struct {
	apiClient *APIClient
}

func NewWSGDPDHCPProfileService(c *APIClient) *WSGDPDHCPProfileService {
	s := new(WSGDPDHCPProfileService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGDPDHCPProfileService() *WSGDPDHCPProfileService {
	return NewWSGDPDHCPProfileService(ss.apiClient)
}

// AddDpDhcpProfiles
//
// Use this API command to create basic DP DHCP profile - basic.
//
// Request Body:
//	 - body *WSGDPProfileDpDhcpProfileBasicBO
func (s *WSGDPDHCPProfileService) AddDpDhcpProfiles(ctx context.Context, body *WSGDPProfileDpDhcpProfileBasicBO) (*WSGDPProfileDpDhcpProfileBasicBO, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGDPProfileDpDhcpProfileBasicBO
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
	req = NewAPIRequest(http.MethodPost, RouteWSGAddDpDhcpProfiles, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGDPProfileDpDhcpProfileBasicBO()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// AddDpDhcpProfilesDpDhcpProfileHostsById
//
// Use this API command to create DP DHCP profile - host.
//
// Request Body:
//	 - body *WSGDPProfileDpDhcpProfileHostBO
//
// Required Parameters:
// - id string
//		- required
func (s *WSGDPDHCPProfileService) AddDpDhcpProfilesDpDhcpProfileHostsById(ctx context.Context, body *WSGDPProfileDpDhcpProfileHostBO, id string) (*WSGDPProfileDpDhcpProfileHostBO, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGDPProfileDpDhcpProfileHostBO
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
	req = NewAPIRequest(http.MethodPost, RouteWSGAddDpDhcpProfilesDpDhcpProfileHostsById, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGDPProfileDpDhcpProfileHostBO()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// AddDpDhcpProfilesDpDhcpProfileOptionSpacesById
//
// Use this API command to create DP DHCP profile - option43 space.
//
// Request Body:
//	 - body *WSGDPProfileDpDhcpProfileOptionSpaceBO
//
// Required Parameters:
// - id string
//		- required
func (s *WSGDPDHCPProfileService) AddDpDhcpProfilesDpDhcpProfileOptionSpacesById(ctx context.Context, body *WSGDPProfileDpDhcpProfileOptionSpaceBO, id string) (interface{}, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodPost, RouteWSGAddDpDhcpProfilesDpDhcpProfileOptionSpacesById, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = new(interface{})
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// AddDpDhcpProfilesDpDhcpProfilePoolsById
//
// Use this API command to create DP DHCP profile - pool.
//
// Request Body:
//	 - body *WSGDPProfileDpDhcpProfilePoolBO
//
// Required Parameters:
// - id string
//		- required
func (s *WSGDPDHCPProfileService) AddDpDhcpProfilesDpDhcpProfilePoolsById(ctx context.Context, body *WSGDPProfileDpDhcpProfilePoolBO, id string) (*WSGDPProfileDpDhcpProfilePoolBO, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGDPProfileDpDhcpProfilePoolBO
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
	req = NewAPIRequest(http.MethodPost, RouteWSGAddDpDhcpProfilesDpDhcpProfilePoolsById, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGDPProfileDpDhcpProfilePoolBO()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// DeleteDpDhcpProfiles
//
// Use this API command to delete DP DHCP profiles.
//
// Request Body:
//	 - body *WSGDPProfileBulkDelete
func (s *WSGDPDHCPProfileService) DeleteDpDhcpProfiles(ctx context.Context, body *WSGDPProfileBulkDelete) (interface{}, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteDpDhcpProfiles, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = new(interface{})
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// DeleteDpDhcpProfilesById
//
// Use this API command to delete DP DHCP profile.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGDPDHCPProfileService) DeleteDpDhcpProfilesById(ctx context.Context, id string) (interface{}, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteDpDhcpProfilesById, true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = new(interface{})
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// DeleteDpDhcpProfilesDpDhcpProfileHostsByHostId
//
// Use this API command to delete DP DHCP profile - host.
//
// Required Parameters:
// - hostId string
//		- required
// - id string
//		- required
func (s *WSGDPDHCPProfileService) DeleteDpDhcpProfilesDpDhcpProfileHostsByHostId(ctx context.Context, hostId string, id string) (interface{}, *APIResponseMeta, error) {
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
	if err = pkgValidator.VarCtx(ctx, hostId, "required"); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteDpDhcpProfilesDpDhcpProfileHostsByHostId, true)
	req.SetPathParameter("hostId", hostId)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = new(interface{})
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// DeleteDpDhcpProfilesDpDhcpProfileHostsById
//
// Use this API command to delete DP DHCP profile - hosts.
//
// Request Body:
//	 - body *WSGDPProfileBulkDelete
//
// Required Parameters:
// - id string
//		- required
func (s *WSGDPDHCPProfileService) DeleteDpDhcpProfilesDpDhcpProfileHostsById(ctx context.Context, body *WSGDPProfileBulkDelete, id string) (interface{}, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteDpDhcpProfilesDpDhcpProfileHostsById, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = new(interface{})
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// DeleteDpDhcpProfilesDpDhcpProfileOptionSpacesById
//
// Use this API command to delete DP DHCP profile - option43 spaces.
//
// Request Body:
//	 - body *WSGDPProfileBulkDelete
//
// Required Parameters:
// - id string
//		- required
func (s *WSGDPDHCPProfileService) DeleteDpDhcpProfilesDpDhcpProfileOptionSpacesById(ctx context.Context, body *WSGDPProfileBulkDelete, id string) (interface{}, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteDpDhcpProfilesDpDhcpProfileOptionSpacesById, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = new(interface{})
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// DeleteDpDhcpProfilesDpDhcpProfileOptionSpacesBySpaceId
//
// Use this API command to delete DP DHCP profile - option43 space.
//
// Required Parameters:
// - id string
//		- required
// - spaceId string
//		- required
func (s *WSGDPDHCPProfileService) DeleteDpDhcpProfilesDpDhcpProfileOptionSpacesBySpaceId(ctx context.Context, id string, spaceId string) (interface{}, *APIResponseMeta, error) {
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
	if err = pkgValidator.VarCtx(ctx, spaceId, "required"); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteDpDhcpProfilesDpDhcpProfileOptionSpacesBySpaceId, true)
	req.SetPathParameter("id", id)
	req.SetPathParameter("spaceId", spaceId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = new(interface{})
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// DeleteDpDhcpProfilesDpDhcpProfilePoolsById
//
// Use this API command to delete DP DHCP profile - pools.
//
// Request Body:
//	 - body *WSGDPProfileBulkDelete
//
// Required Parameters:
// - id string
//		- required
func (s *WSGDPDHCPProfileService) DeleteDpDhcpProfilesDpDhcpProfilePoolsById(ctx context.Context, body *WSGDPProfileBulkDelete, id string) (interface{}, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteDpDhcpProfilesDpDhcpProfilePoolsById, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = new(interface{})
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// DeleteDpDhcpProfilesDpDhcpProfilePoolsByPoolId
//
// Use this API command to delete DP DHCP profile - pool.
//
// Required Parameters:
// - id string
//		- required
// - poolId string
//		- required
func (s *WSGDPDHCPProfileService) DeleteDpDhcpProfilesDpDhcpProfilePoolsByPoolId(ctx context.Context, id string, poolId string) (interface{}, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteDpDhcpProfilesDpDhcpProfilePoolsByPoolId, true)
	req.SetPathParameter("id", id)
	req.SetPathParameter("poolId", poolId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = new(interface{})
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindDpDhcpProfiles
//
// Use this API command to retrieve DP profile - basic list.
func (s *WSGDPDHCPProfileService) FindDpDhcpProfiles(ctx context.Context) (*WSGDPProfileDpDhcpProfileBasicBOList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGDPProfileDpDhcpProfileBasicBOList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindDpDhcpProfiles, true)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGDPProfileDpDhcpProfileBasicBOList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// FindDpDhcpProfilesById
//
// Use this API command to retrieve DP profile - basic.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGDPDHCPProfileService) FindDpDhcpProfilesById(ctx context.Context, id string) (*WSGDPProfileDpDhcpProfileBasicBO, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGDPProfileDpDhcpProfileBasicBO
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindDpDhcpProfilesById, true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGDPProfileDpDhcpProfileBasicBO()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// FindDpDhcpProfilesDpDhcpProfileHostsByHostId
//
// Use this API command to retrieve DP DHCP profile - host.
//
// Required Parameters:
// - hostId string
//		- required
// - id string
//		- required
func (s *WSGDPDHCPProfileService) FindDpDhcpProfilesDpDhcpProfileHostsByHostId(ctx context.Context, hostId string, id string) (*WSGDPProfileDpDhcpProfileHostBO, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGDPProfileDpDhcpProfileHostBO
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, hostId, "required"); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindDpDhcpProfilesDpDhcpProfileHostsByHostId, true)
	req.SetPathParameter("hostId", hostId)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGDPProfileDpDhcpProfileHostBO()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// FindDpDhcpProfilesDpDhcpProfileHostsById
//
// Use this API command to retrieve DP DHCP profile - host list.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGDPDHCPProfileService) FindDpDhcpProfilesDpDhcpProfileHostsById(ctx context.Context, id string) (*WSGDPProfileDpDhcpProfileHostBOList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGDPProfileDpDhcpProfileHostBOList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindDpDhcpProfilesDpDhcpProfileHostsById, true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGDPProfileDpDhcpProfileHostBOList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// FindDpDhcpProfilesDpDhcpProfileOptionSpacesById
//
// Use this API command to retrieve DP DHCP profile - option43 space list.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGDPDHCPProfileService) FindDpDhcpProfilesDpDhcpProfileOptionSpacesById(ctx context.Context, id string) (*WSGDPProfileDpDhcpProfileOptionSpaceApplyToBOList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGDPProfileDpDhcpProfileOptionSpaceApplyToBOList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindDpDhcpProfilesDpDhcpProfileOptionSpacesById, true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGDPProfileDpDhcpProfileOptionSpaceApplyToBOList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// FindDpDhcpProfilesDpDhcpProfileOptionSpacesBySpaceId
//
// Use this API command to retrieve DP DHCP profile - option43 space.
//
// Required Parameters:
// - id string
//		- required
// - spaceId string
//		- required
func (s *WSGDPDHCPProfileService) FindDpDhcpProfilesDpDhcpProfileOptionSpacesBySpaceId(ctx context.Context, id string, spaceId string) (*WSGDPProfileDpDhcpProfileOptionSpaceApplyToBO, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGDPProfileDpDhcpProfileOptionSpaceApplyToBO
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, spaceId, "required"); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindDpDhcpProfilesDpDhcpProfileOptionSpacesBySpaceId, true)
	req.SetPathParameter("id", id)
	req.SetPathParameter("spaceId", spaceId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGDPProfileDpDhcpProfileOptionSpaceApplyToBO()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// FindDpDhcpProfilesDpDhcpProfilePoolsById
//
// Use this API command to retrieve DP DHCP profile - pool list.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGDPDHCPProfileService) FindDpDhcpProfilesDpDhcpProfilePoolsById(ctx context.Context, id string) (*WSGDPProfileDpDhcpProfilePoolBOList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGDPProfileDpDhcpProfilePoolBOList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindDpDhcpProfilesDpDhcpProfilePoolsById, true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGDPProfileDpDhcpProfilePoolBOList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// FindDpDhcpProfilesDpDhcpProfilePoolsByPoolId
//
// Use this API command to retrieve DP DHCP profile - pool.
//
// Required Parameters:
// - id string
//		- required
// - poolId string
//		- required
func (s *WSGDPDHCPProfileService) FindDpDhcpProfilesDpDhcpProfilePoolsByPoolId(ctx context.Context, id string, poolId string) (*WSGDPProfileDpDhcpProfilePoolBO, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGDPProfileDpDhcpProfilePoolBO
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
	req = NewAPIRequest(http.MethodGet, RouteWSGFindDpDhcpProfilesDpDhcpProfilePoolsByPoolId, true)
	req.SetPathParameter("id", id)
	req.SetPathParameter("poolId", poolId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGDPProfileDpDhcpProfilePoolBO()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// UpdateDpDhcpProfilesById
//
// Use this API command to modify DP DHCP profile - basic.
//
// Request Body:
//	 - body *WSGDPProfileDpDhcpProfileBasicBO
//
// Required Parameters:
// - id string
//		- required
func (s *WSGDPDHCPProfileService) UpdateDpDhcpProfilesById(ctx context.Context, body *WSGDPProfileDpDhcpProfileBasicBO, id string) (*WSGDPProfileDpDhcpProfileBasicBO, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGDPProfileDpDhcpProfileBasicBO
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
	req = NewAPIRequest(http.MethodPut, RouteWSGUpdateDpDhcpProfilesById, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGDPProfileDpDhcpProfileBasicBO()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// UpdateDpDhcpProfilesDpDhcpProfileHostsByHostId
//
// Use this API command to modify DP DHCP profile - host.
//
// Request Body:
//	 - body *WSGDPProfileDpDhcpProfileHostBO
//
// Required Parameters:
// - hostId string
//		- required
// - id string
//		- required
func (s *WSGDPDHCPProfileService) UpdateDpDhcpProfilesDpDhcpProfileHostsByHostId(ctx context.Context, body *WSGDPProfileDpDhcpProfileHostBO, hostId string, id string) (*WSGDPProfileDpDhcpProfileHostBO, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGDPProfileDpDhcpProfileHostBO
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
	if err = pkgValidator.VarCtx(ctx, hostId, "required"); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPut, RouteWSGUpdateDpDhcpProfilesDpDhcpProfileHostsByHostId, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetPathParameter("hostId", hostId)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGDPProfileDpDhcpProfileHostBO()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// UpdateDpDhcpProfilesDpDhcpProfileOptionSpacesBySpaceId
//
// Use this API command to update DP DHCP profile - option43 space.
//
// Request Body:
//	 - body *WSGDPProfileDpDhcpProfileOptionSpaceBO
//
// Required Parameters:
// - id string
//		- required
// - spaceId string
//		- required
func (s *WSGDPDHCPProfileService) UpdateDpDhcpProfilesDpDhcpProfileOptionSpacesBySpaceId(ctx context.Context, body *WSGDPProfileDpDhcpProfileOptionSpaceBO, id string, spaceId string) (interface{}, *APIResponseMeta, error) {
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
	if err = pkgValidator.VarCtx(ctx, spaceId, "required"); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPut, RouteWSGUpdateDpDhcpProfilesDpDhcpProfileOptionSpacesBySpaceId, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetPathParameter("id", id)
	req.SetPathParameter("spaceId", spaceId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = new(interface{})
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// UpdateDpDhcpProfilesDpDhcpProfilePoolsByPoolId
//
// Use this API command to modify DP DHCP profile - pool.
//
// Request Body:
//	 - body *WSGDPProfileDpDhcpProfilePoolBO
//
// Required Parameters:
// - id string
//		- required
// - poolId string
//		- required
func (s *WSGDPDHCPProfileService) UpdateDpDhcpProfilesDpDhcpProfilePoolsByPoolId(ctx context.Context, body *WSGDPProfileDpDhcpProfilePoolBO, id string, poolId string) (*WSGDPProfileDpDhcpProfilePoolBO, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGDPProfileDpDhcpProfilePoolBO
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
	req = NewAPIRequest(http.MethodPut, RouteWSGUpdateDpDhcpProfilesDpDhcpProfilePoolsByPoolId, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetPathParameter("id", id)
	req.SetPathParameter("poolId", poolId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGDPProfileDpDhcpProfilePoolBO()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}
