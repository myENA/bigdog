package ruckus

// API Version: v9_0

import (
	"context"
	"net/http"
)

type WSGHotspot20wifioperatorprofileService struct {
	apiClient *VSZClient
}

func NewWSGHotspot20wifioperatorprofileService(c *VSZClient) *WSGHotspot20wifioperatorprofileService {
	s := new(WSGHotspot20wifioperatorprofileService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGHotspot20wifioperatorprofileService() *WSGHotspot20wifioperatorprofileService {
	return NewWSGHotspot20wifioperatorprofileService(ss.apiClient)
}

// AddProfilesHs20Operators
//
// Use this API command to create a new Hotspot 2.0 Wi-Fi operator.
//
// Request Body:
//	 - body *WSGProfileHs20Operator
func (s *WSGHotspot20wifioperatorprofileService) AddProfilesHs20Operators(ctx context.Context, body *WSGProfileHs20Operator) (*WSGCommonCreateResult, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodPost, RouteWSGAddProfilesHs20Operators, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonCreateResult()
	rm, err = handleResponse(req, http.StatusCreated, httpResp, resp, err)
	return resp, rm, err
}

// DeleteProfilesHs20Operators
//
// Use this API command to delete multiple Hotspot 2.0 Wi-Fi operators.
//
// Request Body:
//	 - body *WSGCommonBulkDeleteRequest
func (s *WSGHotspot20wifioperatorprofileService) DeleteProfilesHs20Operators(ctx context.Context, body *WSGCommonBulkDeleteRequest) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteProfilesHs20Operators, true)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteProfilesHs20OperatorsById
//
// Use this API command to delete a Hotspot 2.0 Wi-Fi operator.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGHotspot20wifioperatorprofileService) DeleteProfilesHs20OperatorsById(ctx context.Context, id string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteProfilesHs20OperatorsById, true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteProfilesHs20OperatorsCertificateById
//
// Use this API command to disable certificate of a Hotspot 2.0 Wi-Fi operator.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGHotspot20wifioperatorprofileService) DeleteProfilesHs20OperatorsCertificateById(ctx context.Context, id string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteProfilesHs20OperatorsCertificateById, true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// FindProfilesHs20Operators
//
// Use this API command to retrieve list of Hotspot 2.0 Wi-Fi Operators.
//
// Optional Parameters:
// - index string
//		- nullable
// - listSize string
//		- nullable
func (s *WSGHotspot20wifioperatorprofileService) FindProfilesHs20Operators(ctx context.Context, optionalParams map[string][]string) (*WSGProfileHs20OperatorList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGProfileHs20OperatorList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindProfilesHs20Operators, true)
	if v, ok := optionalParams["index"]; ok && len(v) > 0 {
		req.SetQueryParameter("index", v)
	}
	if v, ok := optionalParams["listSize"]; ok && len(v) > 0 {
		req.SetQueryParameter("listSize", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGProfileHs20OperatorList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindProfilesHs20OperatorsById
//
// Use this API command to retrieve a Hotspot 2.0 Wi-Fi operator.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGHotspot20wifioperatorprofileService) FindProfilesHs20OperatorsById(ctx context.Context, id string) (*WSGProfileHs20Operator, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGProfileHs20Operator
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindProfilesHs20OperatorsById, true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGProfileHs20Operator()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindProfilesHs20OperatorsByQueryCriteria
//
// Query hotspot 2.0 Wi-Fi operators.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGHotspot20wifioperatorprofileService) FindProfilesHs20OperatorsByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet) (*WSGProfileHs20OperatorList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGProfileHs20OperatorList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGFindProfilesHs20OperatorsByQueryCriteria, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGProfileHs20OperatorList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// PartialUpdateProfilesHs20OperatorsById
//
// Use this API command to modify the configuration of a Hotspot 2.0 Wi-Fi operator.
//
// Request Body:
//	 - body *WSGProfileModifyHS20Operator
//
// Required Parameters:
// - id string
//		- required
func (s *WSGHotspot20wifioperatorprofileService) PartialUpdateProfilesHs20OperatorsById(ctx context.Context, body *WSGProfileModifyHS20Operator, id string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateProfilesHs20OperatorsById, true)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// UpdateProfilesHs20OperatorsById
//
// Use this API command to modify entire configuration of a Hotspot 2.0 Wi-Fi operator.
//
// Request Body:
//	 - body *WSGProfileHs20Operator
//
// Required Parameters:
// - id string
//		- required
func (s *WSGHotspot20wifioperatorprofileService) UpdateProfilesHs20OperatorsById(ctx context.Context, body *WSGProfileHs20Operator, id string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPut, RouteWSGUpdateProfilesHs20OperatorsById, true)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}
