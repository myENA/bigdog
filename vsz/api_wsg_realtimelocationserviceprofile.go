package vsz

// API Version: v9_0

import (
	"context"
	"net/http"
)

type WSGRealtimelocationserviceprofileService struct {
	apiClient *APIClient
}

func NewWSGRealtimelocationserviceprofileService(c *APIClient) *WSGRealtimelocationserviceprofileService {
	s := new(WSGRealtimelocationserviceprofileService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGRealtimelocationserviceprofileService() *WSGRealtimelocationserviceprofileService {
	return NewWSGRealtimelocationserviceprofileService(ss.apiClient)
}

// AddRkszonesRealTimeLocationServiceByZoneId
//
// Use this API command to create RTLS Profile.
//
// Request Body:
//	 - body *WSGProfileCreateRtlsProfile
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGRealtimelocationserviceprofileService) AddRkszonesRealTimeLocationServiceByZoneId(ctx context.Context, body *WSGProfileCreateRtlsProfile, zoneId string) (*WSGCommonCreateResult, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodPost, RouteWSGAddRkszonesRealTimeLocationServiceByZoneId, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonCreateResult()
	rm, err = handleResponse(req, http.StatusCreated, httpResp, resp, err)
	return resp, rm, err
}

// DeleteRkszonesRealTimeLocationServiceById
//
// Use this API command to Delete RTLS Profile by profile's ID.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGRealtimelocationserviceprofileService) DeleteRkszonesRealTimeLocationServiceById(ctx context.Context, id string, zoneId string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteRkszonesRealTimeLocationServiceById, true)
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// FindRkszonesRealTimeLocationServiceById
//
// Use this API command to Get RTLS Profile by profile's ID.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGRealtimelocationserviceprofileService) FindRkszonesRealTimeLocationServiceById(ctx context.Context, id string, zoneId string) (*WSGProfileCreateRtlsProfile, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodGet, RouteWSGFindRkszonesRealTimeLocationServiceById, true)
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGProfileCreateRtlsProfile()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindRkszonesRealTimeLocationServiceByZoneId
//
// Use this API command to Get RTLS Profile by zone ID.
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGRealtimelocationserviceprofileService) FindRkszonesRealTimeLocationServiceByZoneId(ctx context.Context, zoneId string) (*WSGProfileRtlsProfileList, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodGet, RouteWSGFindRkszonesRealTimeLocationServiceByZoneId, true)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGProfileRtlsProfileList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// UpdateRkszonesRealTimeLocationServiceById
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
func (s *WSGRealtimelocationserviceprofileService) UpdateRkszonesRealTimeLocationServiceById(ctx context.Context, body *WSGProfileUpdateRtlsProfile, id string, zoneId string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPut, RouteWSGUpdateRkszonesRealTimeLocationServiceById, true)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusOK, httpResp, nil, err)
	return rm, err
}
