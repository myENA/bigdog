package bigdog

// API Version: v9_1

import (
	"context"
	"net/http"
)

type WSGEthernetPortProfileService struct {
	apiClient *VSZClient
}

func NewWSGEthernetPortProfileService(c *VSZClient) *WSGEthernetPortProfileService {
	s := new(WSGEthernetPortProfileService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGEthernetPortProfileService() *WSGEthernetPortProfileService {
	return NewWSGEthernetPortProfileService(ss.apiClient)
}

// AddRkszonesProfileEthernetPortByZoneId
//
// Create a new Ethernet Port Porfile.
//
// Operation ID: addRkszonesProfileEthernetPortByZoneId
// Operation path: /rkszones/{zoneId}/profile/ethernetPort
// Success code: 201 (Created)
//
// Request body:
//	 - body *WSGEthernetPortCreateEthernetPortProfile
//
// Required parameters:
// - zoneId string
//		- required
func (s *WSGEthernetPortProfileService) AddRkszonesProfileEthernetPortByZoneId(ctx context.Context, body *WSGEthernetPortCreateEthernetPortProfile, zoneId string, mutators ...RequestMutator) (*WSGCommonCreateResultAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGCommonCreateResultAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGCommonCreateResultAPIResponse), err
	}
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGAddRkszonesProfileEthernetPortByZoneId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*WSGCommonCreateResultAPIResponse), err
	}
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusCreated, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGCommonCreateResultAPIResponse), err
}

// DeleteRkszonesProfileEthernetPortById
//
// Delete Ethernet Port Porfile.
//
// Operation ID: deleteRkszonesProfileEthernetPortById
// Operation path: /rkszones/{zoneId}/profile/ethernetPort/{id}
// Success code: 204 (No Content)
//
// Required parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGEthernetPortProfileService) DeleteRkszonesProfileEthernetPortById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*EmptyAPIResponse), err
	}
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteRkszonesProfileEthernetPortById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// FindRkszonesProfileEthernetPortById
//
// Retrieve a Ethernet Port Porfile.
//
// Operation ID: findRkszonesProfileEthernetPortById
// Operation path: /rkszones/{zoneId}/profile/ethernetPort/{id}
// Success code: 200 (OK)
//
// Required parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGEthernetPortProfileService) FindRkszonesProfileEthernetPortById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*WSGEthernetPortProfileAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGEthernetPortProfileAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGEthernetPortProfileAPIResponse), err
	}
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindRkszonesProfileEthernetPortById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGEthernetPortProfileAPIResponse), err
}

// FindRkszonesProfileEthernetPortByZoneId
//
// Retrieve a list of Ethernet Port Porfiles within a zone.
//
// Operation ID: findRkszonesProfileEthernetPortByZoneId
// Operation path: /rkszones/{zoneId}/profile/ethernetPort
// Success code: 200 (OK)
//
// Required parameters:
// - zoneId string
//		- required
//
// Optional parameters:
// - index string
//		- nullable
// - listSize string
//		- nullable
func (s *WSGEthernetPortProfileService) FindRkszonesProfileEthernetPortByZoneId(ctx context.Context, zoneId string, optionalParams map[string][]string, mutators ...RequestMutator) (*WSGEthernetPortProfileListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGEthernetPortProfileListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGEthernetPortProfileListAPIResponse), err
	}
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindRkszonesProfileEthernetPortByZoneId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("zoneId", zoneId)
	if v, ok := optionalParams["index"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("index", v)
	}
	if v, ok := optionalParams["listSize"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("listSize", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGEthernetPortProfileListAPIResponse), err
}

// FindServicesEthernetPortProfileByQueryCriteria
//
// Query Ethernet Port Profiles with specified filters.
//
// Operation ID: findServicesEthernetPortProfileByQueryCriteria
// Operation path: /query/services/ethernetPortProfile
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGEthernetPortProfileService) FindServicesEthernetPortProfileByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGFindServicesEthernetPortProfileByQueryCriteria, true)
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

// PartialUpdateRkszonesProfileEthernetPortById
//
// Modify a specific Ethernet Port Porfile.
//
// Operation ID: partialUpdateRkszonesProfileEthernetPortById
// Operation path: /rkszones/{zoneId}/profile/ethernetPort/{id}
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGEthernetPortModifyEthernetPortProfile
//
// Required parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGEthernetPortProfileService) PartialUpdateRkszonesProfileEthernetPortById(ctx context.Context, body *WSGEthernetPortModifyEthernetPortProfile, id string, zoneId string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*EmptyAPIResponse), err
	}
	req = apiRequestFromPool(APISourceVSZ, http.MethodPatch, RouteWSGPartialUpdateRkszonesProfileEthernetPortById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*EmptyAPIResponse), err
	}
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// UpdateRkszonesProfileEthernetPortById
//
// Modify a specific Ethernet Port Porfile.
//
// Operation ID: updateRkszonesProfileEthernetPortById
// Operation path: /rkszones/{zoneId}/profile/ethernetPort/{id}
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGEthernetPortModifyEthernetPortProfile
//
// Required parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGEthernetPortProfileService) UpdateRkszonesProfileEthernetPortById(ctx context.Context, body *WSGEthernetPortModifyEthernetPortProfile, id string, zoneId string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*EmptyAPIResponse), err
	}
	req = apiRequestFromPool(APISourceVSZ, http.MethodPut, RouteWSGUpdateRkszonesProfileEthernetPortById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*EmptyAPIResponse), err
	}
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}
