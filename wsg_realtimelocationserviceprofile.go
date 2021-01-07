package bigdog

// API Version: v9_1

import (
	"context"
	"net/http"
)

type WSGRealTimeLocationServiceProfileService struct {
	apiClient *VSZClient
}

func NewWSGRealTimeLocationServiceProfileService(c *VSZClient) *WSGRealTimeLocationServiceProfileService {
	s := new(WSGRealTimeLocationServiceProfileService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGRealTimeLocationServiceProfileService() *WSGRealTimeLocationServiceProfileService {
	return NewWSGRealTimeLocationServiceProfileService(ss.apiClient)
}

// AddRkszonesRealTimeLocationServiceByZoneId
//
// Operation ID: addRkszonesRealTimeLocationServiceByZoneId
//
// Use this API command to create RTLS Profile.
//
// Request Body:
//	 - body *WSGProfileCreateRtlsProfile
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGRealTimeLocationServiceProfileService) AddRkszonesRealTimeLocationServiceByZoneId(ctx context.Context, body *WSGProfileCreateRtlsProfile, zoneId string, mutators ...RequestMutator) (*WSGCommonCreateResultAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGCommonCreateResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGAddRkszonesRealTimeLocationServiceByZoneId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGCommonCreateResult()
	rm, err = handleAPIResponse(req, http.StatusCreated, httpResp, resp, err)
	return resp, rm, err
}

// DeleteRkszonesRealTimeLocationServiceById
//
// Operation ID: deleteRkszonesRealTimeLocationServiceById
//
// Use this API command to Delete RTLS Profile by profile's ID.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGRealTimeLocationServiceProfileService) DeleteRkszonesRealTimeLocationServiceById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteRkszonesRealTimeLocationServiceById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleAPIResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// FindRkszonesRealTimeLocationServiceById
//
// Operation ID: findRkszonesRealTimeLocationServiceById
//
// Use this API command to Get RTLS Profile by profile's ID.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGRealTimeLocationServiceProfileService) FindRkszonesRealTimeLocationServiceById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*WSGProfileCreateRtlsProfileAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGProfileCreateRtlsProfile
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindRkszonesRealTimeLocationServiceById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGProfileCreateRtlsProfile()
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindRkszonesRealTimeLocationServiceByZoneId
//
// Operation ID: findRkszonesRealTimeLocationServiceByZoneId
//
// Use this API command to Get RTLS Profile by zone ID.
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGRealTimeLocationServiceProfileService) FindRkszonesRealTimeLocationServiceByZoneId(ctx context.Context, zoneId string, mutators ...RequestMutator) (*WSGProfileRtlsProfileListAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGProfileRtlsProfileList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindRkszonesRealTimeLocationServiceByZoneId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGProfileRtlsProfileList()
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// UpdateRkszonesRealTimeLocationServiceById
//
// Operation ID: updateRkszonesRealTimeLocationServiceById
//
// Use this API command to Modify RTLS Profile by profile's ID.
//
// Request Body:
//	 - body *WSGProfileUpdateRtlsProfile
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGRealTimeLocationServiceProfileService) UpdateRkszonesRealTimeLocationServiceById(ctx context.Context, body *WSGProfileUpdateRtlsProfile, id string, zoneId string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = apiRequestFromPool(http.MethodPut, RouteWSGUpdateRkszonesRealTimeLocationServiceById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, nil, err)
	return rm, err
}
