package bigdog

// API Version: v9_0

import (
	"context"
	"net/http"
)

type WSGRestrictedAPAccessProfileService struct {
	apiClient *VSZClient
}

func NewWSGRestrictedAPAccessProfileService(c *VSZClient) *WSGRestrictedAPAccessProfileService {
	s := new(WSGRestrictedAPAccessProfileService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGRestrictedAPAccessProfileService() *WSGRestrictedAPAccessProfileService {
	return NewWSGRestrictedAPAccessProfileService(ss.apiClient)
}

// AddRkszonesRestrictedApAccessProfilesByZoneId
//
// Create a Restricted AP Access Profile.
//
// Request Body:
//	 - body *WSGProfileCreateRestrictedApAccessProfile
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGRestrictedAPAccessProfileService) AddRkszonesRestrictedApAccessProfilesByZoneId(ctx context.Context, body *WSGProfileCreateRestrictedApAccessProfile, zoneId string) (*WSGCommonCreateResult, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodPost, RouteWSGAddRkszonesRestrictedApAccessProfilesByZoneId, true)
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

// DeleteRkszonesRestrictedApAccessProfiles
//
// Use this API command to delete Bulk Restricted AP Access Profile.
//
// Request Body:
//	 - body *WSGCommonBulkDeleteRequest
func (s *WSGRestrictedAPAccessProfileService) DeleteRkszonesRestrictedApAccessProfiles(ctx context.Context, body *WSGCommonBulkDeleteRequest) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteRkszonesRestrictedApAccessProfiles, true)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteRkszonesRestrictedApAccessProfilesById
//
// Delete a Restricted AP Access Profile.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGRestrictedAPAccessProfileService) DeleteRkszonesRestrictedApAccessProfilesById(ctx context.Context, id string, zoneId string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteRkszonesRestrictedApAccessProfilesById, true)
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// FindRkszonesRestrictedApAccessProfilesById
//
// Retrieve a Restricted AP Access Profile.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGRestrictedAPAccessProfileService) FindRkszonesRestrictedApAccessProfilesById(ctx context.Context, id string, zoneId string) (*WSGProfileRestrictedApAccessProfile, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGProfileRestrictedApAccessProfile
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindRkszonesRestrictedApAccessProfilesById, true)
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGProfileRestrictedApAccessProfile()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindRkszonesRestrictedApAccessProfilesByQueryCriteria
//
// Retrieve a list of Restricted AP Access Profile.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGRestrictedAPAccessProfileService) FindRkszonesRestrictedApAccessProfilesByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet) (*WSGProfileRestrictedApAccessProfileArray, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGProfileRestrictedApAccessProfileArray
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGFindRkszonesRestrictedApAccessProfilesByQueryCriteria, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGProfileRestrictedApAccessProfileArray()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindRkszonesRestrictedApAccessProfilesByZoneId
//
// Retrieve Restricted AP Access Profile list.
//
// Required Parameters:
// - zoneId string
//		- required
//
// Optional Parameters:
// - index string
//		- nullable
// - listSize string
//		- nullable
func (s *WSGRestrictedAPAccessProfileService) FindRkszonesRestrictedApAccessProfilesByZoneId(ctx context.Context, zoneId string, optionalParams map[string][]string) (*WSGProfileIdList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGProfileIdList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindRkszonesRestrictedApAccessProfilesByZoneId, true)
	req.SetPathParameter("zoneId", zoneId)
	if v, ok := optionalParams["index"]; ok && len(v) > 0 {
		req.SetQueryParameter("index", v)
	}
	if v, ok := optionalParams["listSize"]; ok && len(v) > 0 {
		req.SetQueryParameter("listSize", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGProfileIdList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// UpdateRkszonesRestrictedApAccessProfilesById
//
// Modify a Restricted AP Access Profile.
//
// Request Body:
//	 - body *WSGProfileModifyRestrictedApAccessProfile
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGRestrictedAPAccessProfileService) UpdateRkszonesRestrictedApAccessProfilesById(ctx context.Context, body *WSGProfileModifyRestrictedApAccessProfile, id string, zoneId string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPut, RouteWSGUpdateRkszonesRestrictedApAccessProfilesById, true)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}
