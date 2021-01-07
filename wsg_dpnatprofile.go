package bigdog

// API Version: v9_1

import (
	"context"
	"net/http"
)

type WSGDPNATProfileService struct {
	apiClient *VSZClient
}

func NewWSGDPNATProfileService(c *VSZClient) *WSGDPNATProfileService {
	s := new(WSGDPNATProfileService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGDPNATProfileService() *WSGDPNATProfileService {
	return NewWSGDPNATProfileService(ss.apiClient)
}

// AddDpNatProfiles
//
// Operation ID: addDpNatProfiles
//
// Use this API command to create DHCP NAT profile - basic.
//
// Request Body:
//	 - body *WSGDPProfileDpNatProfileBasicBO
func (s *WSGDPNATProfileService) AddDpNatProfiles(ctx context.Context, body *WSGDPProfileDpNatProfileBasicBO, mutators ...RequestMutator) (*WSGDPProfileDpNatProfileBasicBOAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodPost, RouteWSGAddDpNatProfiles, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGDPProfileDpNatProfileBasicBO()
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// AddDpNatProfilesDpNatPoolsById
//
// Operation ID: addDpNatProfilesDpNatPoolsById
//
// Use this API command to create DHCP NAT profile - pool.
//
// Request Body:
//	 - body *WSGDPProfileDpNatProfilePoolBO
//
// Required Parameters:
// - id string
//		- required
func (s *WSGDPNATProfileService) AddDpNatProfilesDpNatPoolsById(ctx context.Context, body *WSGDPProfileDpNatProfilePoolBO, id string, mutators ...RequestMutator) (*WSGDPProfileDpNatProfilePoolBOAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodPost, RouteWSGAddDpNatProfilesDpNatPoolsById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGDPProfileDpNatProfilePoolBO()
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// DeleteDpNatProfiles
//
// Operation ID: deleteDpNatProfiles
//
// Use this API command to delete DHCP NAT profiles.
//
// Request Body:
//	 - body *WSGDPProfileBulkDelete
func (s *WSGDPNATProfileService) DeleteDpNatProfiles(ctx context.Context, body *WSGDPProfileBulkDelete, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *RawAPIResponse
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteDpNatProfiles, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = new(RawAPIResponse)
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// DeleteDpNatProfilesById
//
// Operation ID: deleteDpNatProfilesById
//
// Use this API command to delete DHCP NAT profile.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGDPNATProfileService) DeleteDpNatProfilesById(ctx context.Context, id string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *RawAPIResponse
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteDpNatProfilesById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = new(RawAPIResponse)
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// DeleteDpNatProfilesDpNatPoolsById
//
// Operation ID: deleteDpNatProfilesDpNatPoolsById
//
// Use this API command to delete DP NAT profile - pools.
//
// Request Body:
//	 - body *WSGDPProfileBulkDelete
//
// Required Parameters:
// - id string
//		- required
func (s *WSGDPNATProfileService) DeleteDpNatProfilesDpNatPoolsById(ctx context.Context, body *WSGDPProfileBulkDelete, id string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *RawAPIResponse
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteDpNatProfilesDpNatPoolsById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = new(RawAPIResponse)
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// DeleteDpNatProfilesDpNatPoolsByPoolId
//
// Operation ID: deleteDpNatProfilesDpNatPoolsByPoolId
//
// Use this API command to delete DP NAT profile - pool.
//
// Required Parameters:
// - id string
//		- required
// - poolId string
//		- required
func (s *WSGDPNATProfileService) DeleteDpNatProfilesDpNatPoolsByPoolId(ctx context.Context, id string, poolId string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *RawAPIResponse
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteDpNatProfilesDpNatPoolsByPoolId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	req.PathParams.Set("id", id)
	req.PathParams.Set("poolId", poolId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = new(RawAPIResponse)
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindDpNatProfiles
//
// Operation ID: findDpNatProfiles
//
// Use this API command to retrieve DHCP NAT profile - basic list.
func (s *WSGDPNATProfileService) FindDpNatProfiles(ctx context.Context, mutators ...RequestMutator) (*WSGDPProfileDpNatProfileBasicBOListAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindDpNatProfiles, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGDPProfileDpNatProfileBasicBOList()
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindDpNatProfilesById
//
// Operation ID: findDpNatProfilesById
//
// Use this API command to retrieve DHCP NAT profile - basic.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGDPNATProfileService) FindDpNatProfilesById(ctx context.Context, id string, mutators ...RequestMutator) (*WSGDPProfileDpNatProfileBasicBOAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindDpNatProfilesById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGDPProfileDpNatProfileBasicBO()
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindDpNatProfilesDpNatPoolsById
//
// Operation ID: findDpNatProfilesDpNatPoolsById
//
// Use this API command to retrieve DP NAT profile - pool list.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGDPNATProfileService) FindDpNatProfilesDpNatPoolsById(ctx context.Context, id string, mutators ...RequestMutator) (*WSGDPProfileDpNatProfilePoolBOListAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindDpNatProfilesDpNatPoolsById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGDPProfileDpNatProfilePoolBOList()
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindDpNatProfilesDpNatPoolsByPoolId
//
// Operation ID: findDpNatProfilesDpNatPoolsByPoolId
//
// Use this API command to retrieve DP DHCP profile - pool.
//
// Required Parameters:
// - id string
//		- required
// - poolId string
//		- required
func (s *WSGDPNATProfileService) FindDpNatProfilesDpNatPoolsByPoolId(ctx context.Context, id string, poolId string, mutators ...RequestMutator) (*WSGDPProfileDpNatProfilePoolBOAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindDpNatProfilesDpNatPoolsByPoolId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	req.PathParams.Set("id", id)
	req.PathParams.Set("poolId", poolId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGDPProfileDpNatProfilePoolBO()
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// UpdateDpNatProfilesById
//
// Operation ID: updateDpNatProfilesById
//
// Use this API command to modify DHCP NAT profile - basic.
//
// Request Body:
//	 - body *WSGDPProfileDpNatProfileBasicBO
//
// Required Parameters:
// - id string
//		- required
func (s *WSGDPNATProfileService) UpdateDpNatProfilesById(ctx context.Context, body *WSGDPProfileDpNatProfileBasicBO, id string, mutators ...RequestMutator) (*WSGDPProfileDpNatProfileBasicBOAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodPut, RouteWSGUpdateDpNatProfilesById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGDPProfileDpNatProfileBasicBO()
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// UpdateDpNatProfilesDpNatPoolsByPoolId
//
// Operation ID: updateDpNatProfilesDpNatPoolsByPoolId
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
func (s *WSGDPNATProfileService) UpdateDpNatProfilesDpNatPoolsByPoolId(ctx context.Context, body *WSGDPProfileDpNatProfilePoolBO, id string, poolId string, mutators ...RequestMutator) (*WSGDPProfileDpNatProfilePoolBOAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodPut, RouteWSGUpdateDpNatProfilesDpNatPoolsByPoolId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.PathParams.Set("id", id)
	req.PathParams.Set("poolId", poolId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGDPProfileDpNatProfilePoolBO()
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
