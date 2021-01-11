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
// Use this API command to create DHCP NAT profile - basic.
//
// Operation ID: addDpNatProfiles
// Operation path: /dpNatProfiles
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGDPProfileDpNatProfileBasicBO
func (s *WSGDPNATProfileService) AddDpNatProfiles(ctx context.Context, body *WSGDPProfileDpNatProfileBasicBO, mutators ...RequestMutator) (*WSGDPProfileDpNatProfileBasicBOAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGDPProfileDpNatProfileBasicBOAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGDPProfileDpNatProfileBasicBOAPIResponse), err
	}
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGAddDpNatProfiles, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*WSGDPProfileDpNatProfileBasicBOAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGDPProfileDpNatProfileBasicBOAPIResponse), err
}

// AddDpNatProfilesDpNatPoolsById
//
// Use this API command to create DHCP NAT profile - pool.
//
// Operation ID: addDpNatProfilesDpNatPoolsById
// Operation path: /dpNatProfiles/{id}/dpNatPools
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGDPProfileDpNatProfilePoolBO
//
// Required parameters:
// - id string
//		- required
func (s *WSGDPNATProfileService) AddDpNatProfilesDpNatPoolsById(ctx context.Context, body *WSGDPProfileDpNatProfilePoolBO, id string, mutators ...RequestMutator) (*WSGDPProfileDpNatProfilePoolBOAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGDPProfileDpNatProfilePoolBOAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGDPProfileDpNatProfilePoolBOAPIResponse), err
	}
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGAddDpNatProfilesDpNatPoolsById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*WSGDPProfileDpNatProfilePoolBOAPIResponse), err
	}
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGDPProfileDpNatProfilePoolBOAPIResponse), err
}

// DeleteDpNatProfiles
//
// Use this API command to delete DHCP NAT profiles.
//
// Operation ID: deleteDpNatProfiles
// Operation path: /dpNatProfiles
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGDPProfileBulkDelete
func (s *WSGDPNATProfileService) DeleteDpNatProfiles(ctx context.Context, body *WSGDPProfileBulkDelete, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteDpNatProfiles, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*RawAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// DeleteDpNatProfilesById
//
// Use this API command to delete DHCP NAT profile.
//
// Operation ID: deleteDpNatProfilesById
// Operation path: /dpNatProfiles/{id}
// Success code: 200 (OK)
//
// Required parameters:
// - id string
//		- required
func (s *WSGDPNATProfileService) DeleteDpNatProfilesById(ctx context.Context, id string, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteDpNatProfilesById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// DeleteDpNatProfilesDpNatPoolsById
//
// Use this API command to delete DP NAT profile - pools.
//
// Operation ID: deleteDpNatProfilesDpNatPoolsById
// Operation path: /dpNatProfiles/{id}/dpNatPools
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGDPProfileBulkDelete
//
// Required parameters:
// - id string
//		- required
func (s *WSGDPNATProfileService) DeleteDpNatProfilesDpNatPoolsById(ctx context.Context, body *WSGDPProfileBulkDelete, id string, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteDpNatProfilesDpNatPoolsById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// DeleteDpNatProfilesDpNatPoolsByPoolId
//
// Use this API command to delete DP NAT profile - pool.
//
// Operation ID: deleteDpNatProfilesDpNatPoolsByPoolId
// Operation path: /dpNatProfiles/{id}/dpNatPools/{poolId}
// Success code: 200 (OK)
//
// Required parameters:
// - id string
//		- required
// - poolId string
//		- required
func (s *WSGDPNATProfileService) DeleteDpNatProfilesDpNatPoolsByPoolId(ctx context.Context, id string, poolId string, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteDpNatProfilesDpNatPoolsByPoolId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	req.PathParams.Set("poolId", poolId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// FindDpNatProfiles
//
// Use this API command to retrieve DHCP NAT profile - basic list.
//
// Operation ID: findDpNatProfiles
// Operation path: /dpNatProfiles
// Success code: 200 (OK)
func (s *WSGDPNATProfileService) FindDpNatProfiles(ctx context.Context, mutators ...RequestMutator) (*WSGDPProfileDpNatProfileBasicBOListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGDPProfileDpNatProfileBasicBOListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGDPProfileDpNatProfileBasicBOListAPIResponse), err
	}
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindDpNatProfiles, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGDPProfileDpNatProfileBasicBOListAPIResponse), err
}

// FindDpNatProfilesById
//
// Use this API command to retrieve DHCP NAT profile - basic.
//
// Operation ID: findDpNatProfilesById
// Operation path: /dpNatProfiles/{id}
// Success code: 200 (OK)
//
// Required parameters:
// - id string
//		- required
func (s *WSGDPNATProfileService) FindDpNatProfilesById(ctx context.Context, id string, mutators ...RequestMutator) (*WSGDPProfileDpNatProfileBasicBOAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGDPProfileDpNatProfileBasicBOAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGDPProfileDpNatProfileBasicBOAPIResponse), err
	}
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindDpNatProfilesById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGDPProfileDpNatProfileBasicBOAPIResponse), err
}

// FindDpNatProfilesDpNatPoolsById
//
// Use this API command to retrieve DP NAT profile - pool list.
//
// Operation ID: findDpNatProfilesDpNatPoolsById
// Operation path: /dpNatProfiles/{id}/dpNatPools
// Success code: 200 (OK)
//
// Required parameters:
// - id string
//		- required
func (s *WSGDPNATProfileService) FindDpNatProfilesDpNatPoolsById(ctx context.Context, id string, mutators ...RequestMutator) (*WSGDPProfileDpNatProfilePoolBOListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGDPProfileDpNatProfilePoolBOListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGDPProfileDpNatProfilePoolBOListAPIResponse), err
	}
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindDpNatProfilesDpNatPoolsById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGDPProfileDpNatProfilePoolBOListAPIResponse), err
}

// FindDpNatProfilesDpNatPoolsByPoolId
//
// Use this API command to retrieve DP DHCP profile - pool.
//
// Operation ID: findDpNatProfilesDpNatPoolsByPoolId
// Operation path: /dpNatProfiles/{id}/dpNatPools/{poolId}
// Success code: 200 (OK)
//
// Required parameters:
// - id string
//		- required
// - poolId string
//		- required
func (s *WSGDPNATProfileService) FindDpNatProfilesDpNatPoolsByPoolId(ctx context.Context, id string, poolId string, mutators ...RequestMutator) (*WSGDPProfileDpNatProfilePoolBOAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGDPProfileDpNatProfilePoolBOAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGDPProfileDpNatProfilePoolBOAPIResponse), err
	}
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindDpNatProfilesDpNatPoolsByPoolId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	req.PathParams.Set("poolId", poolId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGDPProfileDpNatProfilePoolBOAPIResponse), err
}

// UpdateDpNatProfilesById
//
// Use this API command to modify DHCP NAT profile - basic.
//
// Operation ID: updateDpNatProfilesById
// Operation path: /dpNatProfiles/{id}
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGDPProfileDpNatProfileBasicBO
//
// Required parameters:
// - id string
//		- required
func (s *WSGDPNATProfileService) UpdateDpNatProfilesById(ctx context.Context, body *WSGDPProfileDpNatProfileBasicBO, id string, mutators ...RequestMutator) (*WSGDPProfileDpNatProfileBasicBOAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGDPProfileDpNatProfileBasicBOAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGDPProfileDpNatProfileBasicBOAPIResponse), err
	}
	req = apiRequestFromPool(APISourceVSZ, http.MethodPut, RouteWSGUpdateDpNatProfilesById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*WSGDPProfileDpNatProfileBasicBOAPIResponse), err
	}
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGDPProfileDpNatProfileBasicBOAPIResponse), err
}

// UpdateDpNatProfilesDpNatPoolsByPoolId
//
// Use this API command to modify DHCP NAT profile - pool.
//
// Operation ID: updateDpNatProfilesDpNatPoolsByPoolId
// Operation path: /dpNatProfiles/{id}/dpNatPools/{poolId}
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGDPProfileDpNatProfilePoolBO
//
// Required parameters:
// - id string
//		- required
// - poolId string
//		- required
func (s *WSGDPNATProfileService) UpdateDpNatProfilesDpNatPoolsByPoolId(ctx context.Context, body *WSGDPProfileDpNatProfilePoolBO, id string, poolId string, mutators ...RequestMutator) (*WSGDPProfileDpNatProfilePoolBOAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGDPProfileDpNatProfilePoolBOAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGDPProfileDpNatProfilePoolBOAPIResponse), err
	}
	req = apiRequestFromPool(APISourceVSZ, http.MethodPut, RouteWSGUpdateDpNatProfilesDpNatPoolsByPoolId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*WSGDPProfileDpNatProfilePoolBOAPIResponse), err
	}
	req.PathParams.Set("id", id)
	req.PathParams.Set("poolId", poolId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGDPProfileDpNatProfilePoolBOAPIResponse), err
}
