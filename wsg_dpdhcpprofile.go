package bigdog

// API Version: v9_1

import (
	"context"
	"net/http"
)

type WSGDPDHCPProfileService struct {
	apiClient *VSZClient
}

func NewWSGDPDHCPProfileService(c *VSZClient) *WSGDPDHCPProfileService {
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
// Operation ID: addDpDhcpProfiles
// Operation path: /dpDhcpProfiles
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGDPProfileDpDhcpProfileBasicBO
func (s *WSGDPDHCPProfileService) AddDpDhcpProfiles(ctx context.Context, body *WSGDPProfileDpDhcpProfileBasicBO, mutators ...RequestMutator) (*WSGDPProfileDpDhcpProfileBasicBOAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGDPProfileDpDhcpProfileBasicBOAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGDPProfileDpDhcpProfileBasicBOAPIResponse), err
	}
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGAddDpDhcpProfiles, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*WSGDPProfileDpDhcpProfileBasicBOAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGDPProfileDpDhcpProfileBasicBOAPIResponse), err
}

// AddDpDhcpProfilesDpDhcpProfileHostsById
//
// Use this API command to create DP DHCP profile - host.
//
// Operation ID: addDpDhcpProfilesDpDhcpProfileHostsById
// Operation path: /dpDhcpProfiles/{id}/dpDhcpProfileHosts
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGDPProfileDpDhcpProfileHostBO
//
// Required parameters:
// - id string
//		- required
func (s *WSGDPDHCPProfileService) AddDpDhcpProfilesDpDhcpProfileHostsById(ctx context.Context, body *WSGDPProfileDpDhcpProfileHostBO, id string, mutators ...RequestMutator) (*WSGDPProfileDpDhcpProfileHostBOAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGDPProfileDpDhcpProfileHostBOAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGDPProfileDpDhcpProfileHostBOAPIResponse), err
	}
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGAddDpDhcpProfilesDpDhcpProfileHostsById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*WSGDPProfileDpDhcpProfileHostBOAPIResponse), err
	}
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGDPProfileDpDhcpProfileHostBOAPIResponse), err
}

// AddDpDhcpProfilesDpDhcpProfileOptionSpacesById
//
// Use this API command to create DP DHCP profile - option43 space.
//
// Operation ID: addDpDhcpProfilesDpDhcpProfileOptionSpacesById
// Operation path: /dpDhcpProfiles/{id}/dpDhcpProfileOptionSpaces
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGDPProfileDpDhcpProfileOptionSpaceBO
//
// Required parameters:
// - id string
//		- required
func (s *WSGDPDHCPProfileService) AddDpDhcpProfilesDpDhcpProfileOptionSpacesById(ctx context.Context, body *WSGDPProfileDpDhcpProfileOptionSpaceBO, id string, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGAddDpDhcpProfilesDpDhcpProfileOptionSpacesById, true)
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

// AddDpDhcpProfilesDpDhcpProfilePoolsById
//
// Use this API command to create DP DHCP profile - pool.
//
// Operation ID: addDpDhcpProfilesDpDhcpProfilePoolsById
// Operation path: /dpDhcpProfiles/{id}/dpDhcpProfilePools
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGDPProfileDpDhcpProfilePoolBO
//
// Required parameters:
// - id string
//		- required
func (s *WSGDPDHCPProfileService) AddDpDhcpProfilesDpDhcpProfilePoolsById(ctx context.Context, body *WSGDPProfileDpDhcpProfilePoolBO, id string, mutators ...RequestMutator) (*WSGDPProfileDpDhcpProfilePoolBOAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGDPProfileDpDhcpProfilePoolBOAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGDPProfileDpDhcpProfilePoolBOAPIResponse), err
	}
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGAddDpDhcpProfilesDpDhcpProfilePoolsById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*WSGDPProfileDpDhcpProfilePoolBOAPIResponse), err
	}
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGDPProfileDpDhcpProfilePoolBOAPIResponse), err
}

// DeleteDpDhcpProfiles
//
// Use this API command to delete DP DHCP profiles.
//
// Operation ID: deleteDpDhcpProfiles
// Operation path: /dpDhcpProfiles
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGDPProfileBulkDelete
func (s *WSGDPDHCPProfileService) DeleteDpDhcpProfiles(ctx context.Context, body *WSGDPProfileBulkDelete, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteDpDhcpProfiles, true)
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

// DeleteDpDhcpProfilesById
//
// Use this API command to delete DP DHCP profile.
//
// Operation ID: deleteDpDhcpProfilesById
// Operation path: /dpDhcpProfiles/{id}
// Success code: 200 (OK)
//
// Required parameters:
// - id string
//		- required
func (s *WSGDPDHCPProfileService) DeleteDpDhcpProfilesById(ctx context.Context, id string, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteDpDhcpProfilesById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// DeleteDpDhcpProfilesDpDhcpProfileHostsByHostId
//
// Use this API command to delete DP DHCP profile - host.
//
// Operation ID: deleteDpDhcpProfilesDpDhcpProfileHostsByHostId
// Operation path: /dpDhcpProfiles/{id}/dpDhcpProfileHosts/{hostId}
// Success code: 200 (OK)
//
// Required parameters:
// - hostId string
//		- required
// - id string
//		- required
func (s *WSGDPDHCPProfileService) DeleteDpDhcpProfilesDpDhcpProfileHostsByHostId(ctx context.Context, hostId string, id string, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteDpDhcpProfilesDpDhcpProfileHostsByHostId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("hostId", hostId)
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// DeleteDpDhcpProfilesDpDhcpProfileHostsById
//
// Use this API command to delete DP DHCP profile - hosts.
//
// Operation ID: deleteDpDhcpProfilesDpDhcpProfileHostsById
// Operation path: /dpDhcpProfiles/{id}/dpDhcpProfileHosts
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGDPProfileBulkDelete
//
// Required parameters:
// - id string
//		- required
func (s *WSGDPDHCPProfileService) DeleteDpDhcpProfilesDpDhcpProfileHostsById(ctx context.Context, body *WSGDPProfileBulkDelete, id string, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteDpDhcpProfilesDpDhcpProfileHostsById, true)
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

// DeleteDpDhcpProfilesDpDhcpProfileOptionSpacesById
//
// Use this API command to delete DP DHCP profile - option43 spaces.
//
// Operation ID: deleteDpDhcpProfilesDpDhcpProfileOptionSpacesById
// Operation path: /dpDhcpProfiles/{id}/dpDhcpProfileOptionSpaces
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGDPProfileBulkDelete
//
// Required parameters:
// - id string
//		- required
func (s *WSGDPDHCPProfileService) DeleteDpDhcpProfilesDpDhcpProfileOptionSpacesById(ctx context.Context, body *WSGDPProfileBulkDelete, id string, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteDpDhcpProfilesDpDhcpProfileOptionSpacesById, true)
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

// DeleteDpDhcpProfilesDpDhcpProfileOptionSpacesBySpaceId
//
// Use this API command to delete DP DHCP profile - option43 space.
//
// Operation ID: deleteDpDhcpProfilesDpDhcpProfileOptionSpacesBySpaceId
// Operation path: /dpDhcpProfiles/{id}/dpDhcpProfileOptionSpaces/{spaceId}
// Success code: 200 (OK)
//
// Required parameters:
// - id string
//		- required
// - spaceId string
//		- required
func (s *WSGDPDHCPProfileService) DeleteDpDhcpProfilesDpDhcpProfileOptionSpacesBySpaceId(ctx context.Context, id string, spaceId string, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteDpDhcpProfilesDpDhcpProfileOptionSpacesBySpaceId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	req.PathParams.Set("spaceId", spaceId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// DeleteDpDhcpProfilesDpDhcpProfilePoolsById
//
// Use this API command to delete DP DHCP profile - pools.
//
// Operation ID: deleteDpDhcpProfilesDpDhcpProfilePoolsById
// Operation path: /dpDhcpProfiles/{id}/dpDhcpProfilePools
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGDPProfileBulkDelete
//
// Required parameters:
// - id string
//		- required
func (s *WSGDPDHCPProfileService) DeleteDpDhcpProfilesDpDhcpProfilePoolsById(ctx context.Context, body *WSGDPProfileBulkDelete, id string, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteDpDhcpProfilesDpDhcpProfilePoolsById, true)
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

// DeleteDpDhcpProfilesDpDhcpProfilePoolsByPoolId
//
// Use this API command to delete DP DHCP profile - pool.
//
// Operation ID: deleteDpDhcpProfilesDpDhcpProfilePoolsByPoolId
// Operation path: /dpDhcpProfiles/{id}/dpDhcpProfilePools/{poolId}
// Success code: 200 (OK)
//
// Required parameters:
// - id string
//		- required
// - poolId string
//		- required
func (s *WSGDPDHCPProfileService) DeleteDpDhcpProfilesDpDhcpProfilePoolsByPoolId(ctx context.Context, id string, poolId string, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteDpDhcpProfilesDpDhcpProfilePoolsByPoolId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	req.PathParams.Set("poolId", poolId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// FindDpDhcpProfiles
//
// Use this API command to retrieve DP profile - basic list.
//
// Operation ID: findDpDhcpProfiles
// Operation path: /dpDhcpProfiles
// Success code: 200 (OK)
func (s *WSGDPDHCPProfileService) FindDpDhcpProfiles(ctx context.Context, mutators ...RequestMutator) (*WSGDPProfileDpDhcpProfileBasicBOListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGDPProfileDpDhcpProfileBasicBOListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGDPProfileDpDhcpProfileBasicBOListAPIResponse), err
	}
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindDpDhcpProfiles, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGDPProfileDpDhcpProfileBasicBOListAPIResponse), err
}

// FindDpDhcpProfilesById
//
// Use this API command to retrieve DP profile - basic.
//
// Operation ID: findDpDhcpProfilesById
// Operation path: /dpDhcpProfiles/{id}
// Success code: 200 (OK)
//
// Required parameters:
// - id string
//		- required
func (s *WSGDPDHCPProfileService) FindDpDhcpProfilesById(ctx context.Context, id string, mutators ...RequestMutator) (*WSGDPProfileDpDhcpProfileBasicBOAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGDPProfileDpDhcpProfileBasicBOAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGDPProfileDpDhcpProfileBasicBOAPIResponse), err
	}
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindDpDhcpProfilesById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGDPProfileDpDhcpProfileBasicBOAPIResponse), err
}

// FindDpDhcpProfilesDpDhcpProfileHostsByHostId
//
// Use this API command to retrieve DP DHCP profile - host.
//
// Operation ID: findDpDhcpProfilesDpDhcpProfileHostsByHostId
// Operation path: /dpDhcpProfiles/{id}/dpDhcpProfileHosts/{hostId}
// Success code: 200 (OK)
//
// Required parameters:
// - hostId string
//		- required
// - id string
//		- required
func (s *WSGDPDHCPProfileService) FindDpDhcpProfilesDpDhcpProfileHostsByHostId(ctx context.Context, hostId string, id string, mutators ...RequestMutator) (*WSGDPProfileDpDhcpProfileHostBOAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGDPProfileDpDhcpProfileHostBOAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGDPProfileDpDhcpProfileHostBOAPIResponse), err
	}
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindDpDhcpProfilesDpDhcpProfileHostsByHostId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("hostId", hostId)
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGDPProfileDpDhcpProfileHostBOAPIResponse), err
}

// FindDpDhcpProfilesDpDhcpProfileHostsById
//
// Use this API command to retrieve DP DHCP profile - host list.
//
// Operation ID: findDpDhcpProfilesDpDhcpProfileHostsById
// Operation path: /dpDhcpProfiles/{id}/dpDhcpProfileHosts
// Success code: 200 (OK)
//
// Required parameters:
// - id string
//		- required
func (s *WSGDPDHCPProfileService) FindDpDhcpProfilesDpDhcpProfileHostsById(ctx context.Context, id string, mutators ...RequestMutator) (*WSGDPProfileDpDhcpProfileHostBOListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGDPProfileDpDhcpProfileHostBOListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGDPProfileDpDhcpProfileHostBOListAPIResponse), err
	}
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindDpDhcpProfilesDpDhcpProfileHostsById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGDPProfileDpDhcpProfileHostBOListAPIResponse), err
}

// FindDpDhcpProfilesDpDhcpProfileOptionSpacesById
//
// Use this API command to retrieve DP DHCP profile - option43 space list.
//
// Operation ID: findDpDhcpProfilesDpDhcpProfileOptionSpacesById
// Operation path: /dpDhcpProfiles/{id}/dpDhcpProfileOptionSpaces
// Success code: 200 (OK)
//
// Required parameters:
// - id string
//		- required
func (s *WSGDPDHCPProfileService) FindDpDhcpProfilesDpDhcpProfileOptionSpacesById(ctx context.Context, id string, mutators ...RequestMutator) (*WSGDPProfileDpDhcpProfileOptionSpaceApplyToBOListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGDPProfileDpDhcpProfileOptionSpaceApplyToBOListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGDPProfileDpDhcpProfileOptionSpaceApplyToBOListAPIResponse), err
	}
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindDpDhcpProfilesDpDhcpProfileOptionSpacesById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGDPProfileDpDhcpProfileOptionSpaceApplyToBOListAPIResponse), err
}

// FindDpDhcpProfilesDpDhcpProfileOptionSpacesBySpaceId
//
// Use this API command to retrieve DP DHCP profile - option43 space.
//
// Operation ID: findDpDhcpProfilesDpDhcpProfileOptionSpacesBySpaceId
// Operation path: /dpDhcpProfiles/{id}/dpDhcpProfileOptionSpaces/{spaceId}
// Success code: 200 (OK)
//
// Required parameters:
// - id string
//		- required
// - spaceId string
//		- required
func (s *WSGDPDHCPProfileService) FindDpDhcpProfilesDpDhcpProfileOptionSpacesBySpaceId(ctx context.Context, id string, spaceId string, mutators ...RequestMutator) (*WSGDPProfileDpDhcpProfileOptionSpaceApplyToBOAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGDPProfileDpDhcpProfileOptionSpaceApplyToBOAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGDPProfileDpDhcpProfileOptionSpaceApplyToBOAPIResponse), err
	}
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindDpDhcpProfilesDpDhcpProfileOptionSpacesBySpaceId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	req.PathParams.Set("spaceId", spaceId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGDPProfileDpDhcpProfileOptionSpaceApplyToBOAPIResponse), err
}

// FindDpDhcpProfilesDpDhcpProfilePoolsById
//
// Use this API command to retrieve DP DHCP profile - pool list.
//
// Operation ID: findDpDhcpProfilesDpDhcpProfilePoolsById
// Operation path: /dpDhcpProfiles/{id}/dpDhcpProfilePools
// Success code: 200 (OK)
//
// Required parameters:
// - id string
//		- required
func (s *WSGDPDHCPProfileService) FindDpDhcpProfilesDpDhcpProfilePoolsById(ctx context.Context, id string, mutators ...RequestMutator) (*WSGDPProfileDpDhcpProfilePoolBOListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGDPProfileDpDhcpProfilePoolBOListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGDPProfileDpDhcpProfilePoolBOListAPIResponse), err
	}
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindDpDhcpProfilesDpDhcpProfilePoolsById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGDPProfileDpDhcpProfilePoolBOListAPIResponse), err
}

// FindDpDhcpProfilesDpDhcpProfilePoolsByPoolId
//
// Use this API command to retrieve DP DHCP profile - pool.
//
// Operation ID: findDpDhcpProfilesDpDhcpProfilePoolsByPoolId
// Operation path: /dpDhcpProfiles/{id}/dpDhcpProfilePools/{poolId}
// Success code: 200 (OK)
//
// Required parameters:
// - id string
//		- required
// - poolId string
//		- required
func (s *WSGDPDHCPProfileService) FindDpDhcpProfilesDpDhcpProfilePoolsByPoolId(ctx context.Context, id string, poolId string, mutators ...RequestMutator) (*WSGDPProfileDpDhcpProfilePoolBOAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGDPProfileDpDhcpProfilePoolBOAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGDPProfileDpDhcpProfilePoolBOAPIResponse), err
	}
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindDpDhcpProfilesDpDhcpProfilePoolsByPoolId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	req.PathParams.Set("poolId", poolId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGDPProfileDpDhcpProfilePoolBOAPIResponse), err
}

// UpdateDpDhcpProfilesById
//
// Use this API command to modify DP DHCP profile - basic.
//
// Operation ID: updateDpDhcpProfilesById
// Operation path: /dpDhcpProfiles/{id}
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGDPProfileDpDhcpProfileBasicBO
//
// Required parameters:
// - id string
//		- required
func (s *WSGDPDHCPProfileService) UpdateDpDhcpProfilesById(ctx context.Context, body *WSGDPProfileDpDhcpProfileBasicBO, id string, mutators ...RequestMutator) (*WSGDPProfileDpDhcpProfileBasicBOAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGDPProfileDpDhcpProfileBasicBOAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGDPProfileDpDhcpProfileBasicBOAPIResponse), err
	}
	req = apiRequestFromPool(APISourceVSZ, http.MethodPut, RouteWSGUpdateDpDhcpProfilesById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*WSGDPProfileDpDhcpProfileBasicBOAPIResponse), err
	}
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGDPProfileDpDhcpProfileBasicBOAPIResponse), err
}

// UpdateDpDhcpProfilesDpDhcpProfileHostsByHostId
//
// Use this API command to modify DP DHCP profile - host.
//
// Operation ID: updateDpDhcpProfilesDpDhcpProfileHostsByHostId
// Operation path: /dpDhcpProfiles/{id}/dpDhcpProfileHosts/{hostId}
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGDPProfileDpDhcpProfileHostBO
//
// Required parameters:
// - hostId string
//		- required
// - id string
//		- required
func (s *WSGDPDHCPProfileService) UpdateDpDhcpProfilesDpDhcpProfileHostsByHostId(ctx context.Context, body *WSGDPProfileDpDhcpProfileHostBO, hostId string, id string, mutators ...RequestMutator) (*WSGDPProfileDpDhcpProfileHostBOAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGDPProfileDpDhcpProfileHostBOAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGDPProfileDpDhcpProfileHostBOAPIResponse), err
	}
	req = apiRequestFromPool(APISourceVSZ, http.MethodPut, RouteWSGUpdateDpDhcpProfilesDpDhcpProfileHostsByHostId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*WSGDPProfileDpDhcpProfileHostBOAPIResponse), err
	}
	req.PathParams.Set("hostId", hostId)
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGDPProfileDpDhcpProfileHostBOAPIResponse), err
}

// UpdateDpDhcpProfilesDpDhcpProfileOptionSpacesBySpaceId
//
// Use this API command to update DP DHCP profile - option43 space.
//
// Operation ID: updateDpDhcpProfilesDpDhcpProfileOptionSpacesBySpaceId
// Operation path: /dpDhcpProfiles/{id}/dpDhcpProfileOptionSpaces/{spaceId}
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGDPProfileDpDhcpProfileOptionSpaceBO
//
// Required parameters:
// - id string
//		- required
// - spaceId string
//		- required
func (s *WSGDPDHCPProfileService) UpdateDpDhcpProfilesDpDhcpProfileOptionSpacesBySpaceId(ctx context.Context, body *WSGDPProfileDpDhcpProfileOptionSpaceBO, id string, spaceId string, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(APISourceVSZ, http.MethodPut, RouteWSGUpdateDpDhcpProfilesDpDhcpProfileOptionSpacesBySpaceId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req.PathParams.Set("id", id)
	req.PathParams.Set("spaceId", spaceId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// UpdateDpDhcpProfilesDpDhcpProfilePoolsByPoolId
//
// Use this API command to modify DP DHCP profile - pool.
//
// Operation ID: updateDpDhcpProfilesDpDhcpProfilePoolsByPoolId
// Operation path: /dpDhcpProfiles/{id}/dpDhcpProfilePools/{poolId}
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGDPProfileDpDhcpProfilePoolBO
//
// Required parameters:
// - id string
//		- required
// - poolId string
//		- required
func (s *WSGDPDHCPProfileService) UpdateDpDhcpProfilesDpDhcpProfilePoolsByPoolId(ctx context.Context, body *WSGDPProfileDpDhcpProfilePoolBO, id string, poolId string, mutators ...RequestMutator) (*WSGDPProfileDpDhcpProfilePoolBOAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGDPProfileDpDhcpProfilePoolBOAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGDPProfileDpDhcpProfilePoolBOAPIResponse), err
	}
	req = apiRequestFromPool(APISourceVSZ, http.MethodPut, RouteWSGUpdateDpDhcpProfilesDpDhcpProfilePoolsByPoolId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*WSGDPProfileDpDhcpProfilePoolBOAPIResponse), err
	}
	req.PathParams.Set("id", id)
	req.PathParams.Set("poolId", poolId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGDPProfileDpDhcpProfilePoolBOAPIResponse), err
}
