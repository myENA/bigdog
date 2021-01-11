package bigdog

// API Version: v9_1

import (
	"context"
	"net/http"
)

type WSGWechatService struct {
	apiClient *VSZClient
}

func NewWSGWechatService(c *VSZClient) *WSGWechatService {
	s := new(WSGWechatService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGWechatService() *WSGWechatService {
	return NewWSGWechatService(ss.apiClient)
}

// AddRkszonesPortalsWechatByZoneId
//
// Use this API command to create wechat portal.
//
// Operation ID: addRkszonesPortalsWechatByZoneId
// Operation path: /rkszones/{zoneId}/portals/wechat
// Success code: 201 (Created)
//
// Request body:
//	 - body *WSGPortalServiceCreateWechat
//
// Required parameters:
// - zoneId string
//		- required
func (s *WSGWechatService) AddRkszonesPortalsWechatByZoneId(ctx context.Context, body *WSGPortalServiceCreateWechat, zoneId string, mutators ...RequestMutator) (*WSGCommonCreateResultAPIResponse, error) {
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
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGAddRkszonesPortalsWechatByZoneId, true)
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

// DeleteRkszonesPortalsWechatById
//
// Use this API command to delete wechat portal.
//
// Operation ID: deleteRkszonesPortalsWechatById
// Operation path: /rkszones/{zoneId}/portals/wechat/{id}
// Success code: 204 (No Content)
//
// Required parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGWechatService) DeleteRkszonesPortalsWechatById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
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
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteRkszonesPortalsWechatById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// FindRkszonesPortalsWechatById
//
// Use this API command to retrieve wechat portal by ID.
//
// Operation ID: findRkszonesPortalsWechatById
// Operation path: /rkszones/{zoneId}/portals/wechat/{id}
// Success code: 200 (OK)
//
// Required parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGWechatService) FindRkszonesPortalsWechatById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*WSGPortalServiceWechatConfigurationAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGPortalServiceWechatConfigurationAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGPortalServiceWechatConfigurationAPIResponse), err
	}
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindRkszonesPortalsWechatById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGPortalServiceWechatConfigurationAPIResponse), err
}

// FindRkszonesPortalsWechatByZoneId
//
// Use this API command to retrieve a list of wechat portal.
//
// Operation ID: findRkszonesPortalsWechatByZoneId
// Operation path: /rkszones/{zoneId}/portals/wechat
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
func (s *WSGWechatService) FindRkszonesPortalsWechatByZoneId(ctx context.Context, zoneId string, optionalParams map[string][]string, mutators ...RequestMutator) (*WSGPortalServiceListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGPortalServiceListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGPortalServiceListAPIResponse), err
	}
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindRkszonesPortalsWechatByZoneId, true)
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
	return resp.(*WSGPortalServiceListAPIResponse), err
}

// FindServicesWechatProfileByQueryCriteria
//
// Query Wechat Portals with specified filters.
//
// Operation ID: findServicesWechatProfileByQueryCriteria
// Operation path: /query/services/wechatProfile
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGWechatService) FindServicesWechatProfileByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGFindServicesWechatProfileByQueryCriteria, true)
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

// PartialUpdateRkszonesPortalsWechatById
//
// Use this API command to modify the configuration of wechat portal.
//
// Operation ID: partialUpdateRkszonesPortalsWechatById
// Operation path: /rkszones/{zoneId}/portals/wechat/{id}
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGPortalServiceModifyWechat
//
// Required parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGWechatService) PartialUpdateRkszonesPortalsWechatById(ctx context.Context, body *WSGPortalServiceModifyWechat, id string, zoneId string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
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
	req = apiRequestFromPool(APISourceVSZ, http.MethodPatch, RouteWSGPartialUpdateRkszonesPortalsWechatById, true)
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
