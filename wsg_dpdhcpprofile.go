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
// Operation ID: addDpDhcpProfiles
//
// Use this API command to create basic DP DHCP profile - basic.
//
// Request Body:
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
	req = apiRequestFromPool(http.MethodPost, RouteWSGAddDpDhcpProfiles, true)
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
// Operation ID: addDpDhcpProfilesDpDhcpProfileHostsById
//
// Use this API command to create DP DHCP profile - host.
//
// Request Body:
//	 - body *WSGDPProfileDpDhcpProfileHostBO
//
// Required Parameters:
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
	req = apiRequestFromPool(http.MethodPost, RouteWSGAddDpDhcpProfilesDpDhcpProfileHostsById, true)
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
// Operation ID: addDpDhcpProfilesDpDhcpProfileOptionSpacesById
//
// Use this API command to create DP DHCP profile - option43 space.
//
// Request Body:
//	 - body *WSGDPProfileDpDhcpProfileOptionSpaceBO
//
// Required Parameters:
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
	req = apiRequestFromPool(http.MethodPost, RouteWSGAddDpDhcpProfilesDpDhcpProfileOptionSpacesById, true)
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
// Operation ID: addDpDhcpProfilesDpDhcpProfilePoolsById
//
// Use this API command to create DP DHCP profile - pool.
//
// Request Body:
//	 - body *WSGDPProfileDpDhcpProfilePoolBO
//
// Required Parameters:
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
	req = apiRequestFromPool(http.MethodPost, RouteWSGAddDpDhcpProfilesDpDhcpProfilePoolsById, true)
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
// Operation ID: deleteDpDhcpProfiles
//
// Use this API command to delete DP DHCP profiles.
//
// Request Body:
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
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteDpDhcpProfiles, true)
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
// Operation ID: deleteDpDhcpProfilesById
//
// Use this API command to delete DP DHCP profile.
//
// Required Parameters:
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
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteDpDhcpProfilesById, true)
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
// Operation ID: deleteDpDhcpProfilesDpDhcpProfileHostsByHostId
//
// Use this API command to delete DP DHCP profile - host.
//
// Required Parameters:
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
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteDpDhcpProfilesDpDhcpProfileHostsByHostId, true)
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
// Operation ID: deleteDpDhcpProfilesDpDhcpProfileHostsById
//
// Use this API command to delete DP DHCP profile - hosts.
//
// Request Body:
//	 - body *WSGDPProfileBulkDelete
//
// Required Parameters:
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
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteDpDhcpProfilesDpDhcpProfileHostsById, true)
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
// Operation ID: deleteDpDhcpProfilesDpDhcpProfileOptionSpacesById
//
// Use this API command to delete DP DHCP profile - option43 spaces.
//
// Request Body:
//	 - body *WSGDPProfileBulkDelete
//
// Required Parameters:
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
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteDpDhcpProfilesDpDhcpProfileOptionSpacesById, true)
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
// Operation ID: deleteDpDhcpProfilesDpDhcpProfileOptionSpacesBySpaceId
//
// Use this API command to delete DP DHCP profile - option43 space.
//
// Required Parameters:
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
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteDpDhcpProfilesDpDhcpProfileOptionSpacesBySpaceId, true)
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
// Operation ID: deleteDpDhcpProfilesDpDhcpProfilePoolsById
//
// Use this API command to delete DP DHCP profile - pools.
//
// Request Body:
//	 - body *WSGDPProfileBulkDelete
//
// Required Parameters:
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
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteDpDhcpProfilesDpDhcpProfilePoolsById, true)
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
// Operation ID: deleteDpDhcpProfilesDpDhcpProfilePoolsByPoolId
//
// Use this API command to delete DP DHCP profile - pool.
//
// Required Parameters:
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
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteDpDhcpProfilesDpDhcpProfilePoolsByPoolId, true)
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
// Operation ID: findDpDhcpProfiles
//
// Use this API command to retrieve DP profile - basic list.
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
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindDpDhcpProfiles, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGDPProfileDpDhcpProfileBasicBOListAPIResponse), err
}

// FindDpDhcpProfilesById
//
// Operation ID: findDpDhcpProfilesById
//
// Use this API command to retrieve DP profile - basic.
//
// Required Parameters:
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
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindDpDhcpProfilesById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGDPProfileDpDhcpProfileBasicBOAPIResponse), err
}

// FindDpDhcpProfilesDpDhcpProfileHostsByHostId
//
// Operation ID: findDpDhcpProfilesDpDhcpProfileHostsByHostId
//
// Use this API command to retrieve DP DHCP profile - host.
//
// Required Parameters:
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
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindDpDhcpProfilesDpDhcpProfileHostsByHostId, true)
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
// Operation ID: findDpDhcpProfilesDpDhcpProfileHostsById
//
// Use this API command to retrieve DP DHCP profile - host list.
//
// Required Parameters:
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
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindDpDhcpProfilesDpDhcpProfileHostsById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGDPProfileDpDhcpProfileHostBOListAPIResponse), err
}

// FindDpDhcpProfilesDpDhcpProfileOptionSpacesById
//
// Operation ID: findDpDhcpProfilesDpDhcpProfileOptionSpacesById
//
// Use this API command to retrieve DP DHCP profile - option43 space list.
//
// Required Parameters:
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
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindDpDhcpProfilesDpDhcpProfileOptionSpacesById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGDPProfileDpDhcpProfileOptionSpaceApplyToBOListAPIResponse), err
}

// FindDpDhcpProfilesDpDhcpProfileOptionSpacesBySpaceId
//
// Operation ID: findDpDhcpProfilesDpDhcpProfileOptionSpacesBySpaceId
//
// Use this API command to retrieve DP DHCP profile - option43 space.
//
// Required Parameters:
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
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindDpDhcpProfilesDpDhcpProfileOptionSpacesBySpaceId, true)
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
// Operation ID: findDpDhcpProfilesDpDhcpProfilePoolsById
//
// Use this API command to retrieve DP DHCP profile - pool list.
//
// Required Parameters:
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
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindDpDhcpProfilesDpDhcpProfilePoolsById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGDPProfileDpDhcpProfilePoolBOListAPIResponse), err
}

// FindDpDhcpProfilesDpDhcpProfilePoolsByPoolId
//
// Operation ID: findDpDhcpProfilesDpDhcpProfilePoolsByPoolId
//
// Use this API command to retrieve DP DHCP profile - pool.
//
// Required Parameters:
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
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindDpDhcpProfilesDpDhcpProfilePoolsByPoolId, true)
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
// Operation ID: updateDpDhcpProfilesById
//
// Use this API command to modify DP DHCP profile - basic.
//
// Request Body:
//	 - body *WSGDPProfileDpDhcpProfileBasicBO
//
// Required Parameters:
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
	req = apiRequestFromPool(http.MethodPut, RouteWSGUpdateDpDhcpProfilesById, true)
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
// Operation ID: updateDpDhcpProfilesDpDhcpProfileHostsByHostId
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
	req = apiRequestFromPool(http.MethodPut, RouteWSGUpdateDpDhcpProfilesDpDhcpProfileHostsByHostId, true)
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
// Operation ID: updateDpDhcpProfilesDpDhcpProfileOptionSpacesBySpaceId
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
	req = apiRequestFromPool(http.MethodPut, RouteWSGUpdateDpDhcpProfilesDpDhcpProfileOptionSpacesBySpaceId, true)
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
// Operation ID: updateDpDhcpProfilesDpDhcpProfilePoolsByPoolId
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
	req = apiRequestFromPool(http.MethodPut, RouteWSGUpdateDpDhcpProfilesDpDhcpProfilePoolsByPoolId, true)
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
