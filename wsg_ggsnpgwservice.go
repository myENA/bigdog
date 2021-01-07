package bigdog

// API Version: v9_1

import (
	"context"
	"net/http"
)

type WSGGGSNPGWServiceService struct {
	apiClient *VSZClient
}

func NewWSGGGSNPGWServiceService(c *VSZClient) *WSGGGSNPGWServiceService {
	s := new(WSGGGSNPGWServiceService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGGGSNPGWServiceService() *WSGGGSNPGWServiceService {
	return NewWSGGGSNPGWServiceService(ss.apiClient)
}

// DeleteServicesGgsnDnsServerList
//
// Operation ID: deleteServicesGgsnDnsServerList
//
// Use this API command to Disable the dns server list of GGSN/PGW.
func (s *WSGGGSNPGWServiceService) DeleteServicesGgsnDnsServerList(ctx context.Context, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteServicesGgsnDnsServerList, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleAPIResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteServicesGgsnGgsnList
//
// Operation ID: deleteServicesGgsnGgsnList
//
// Use this API command to disable the ggsn server list of GGSN/PGW.
func (s *WSGGGSNPGWServiceService) DeleteServicesGgsnGgsnList(ctx context.Context, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteServicesGgsnGgsnList, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleAPIResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// FindServicesGgsn
//
// Operation ID: findServicesGgsn
//
// Use this API command to retrieve GGSN/PGW setting.
func (s *WSGGGSNPGWServiceService) FindServicesGgsn(ctx context.Context, mutators ...RequestMutator) (*WSGServiceGgsnConfigAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGServiceGgsnConfig
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindServicesGgsn, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGServiceGgsnConfig()
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// PartialUpdateServicesGgsn
//
// Operation ID: partialUpdateServicesGgsn
//
// Use this API command to modify the setting of GGSN/PGW.
//
// Request Body:
//	 - body *WSGServiceGgsnConfig
func (s *WSGGGSNPGWServiceService) PartialUpdateServicesGgsn(ctx context.Context, body *WSGServiceGgsnConfig, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = apiRequestFromPool(http.MethodPatch, RouteWSGPartialUpdateServicesGgsn, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleAPIResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// PartialUpdateServicesGgsnDnsServerList
//
// Operation ID: partialUpdateServicesGgsnDnsServerList
//
// Use this API command to modify the dns server list of GGSN/PGW.
//
// Request Body:
//	 - body WSGServiceDnsServerList
func (s *WSGGGSNPGWServiceService) PartialUpdateServicesGgsnDnsServerList(ctx context.Context, body WSGServiceDnsServerList, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = apiRequestFromPool(http.MethodPatch, RouteWSGPartialUpdateServicesGgsnDnsServerList, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleAPIResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// PartialUpdateServicesGgsnGgsnList
//
// Operation ID: partialUpdateServicesGgsnGgsnList
//
// Use this API command to modify the ggsn server list of GGSN/PGW.
//
// Request Body:
//	 - body WSGServiceGgsnList
func (s *WSGGGSNPGWServiceService) PartialUpdateServicesGgsnGgsnList(ctx context.Context, body WSGServiceGgsnList, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = apiRequestFromPool(http.MethodPatch, RouteWSGPartialUpdateServicesGgsnGgsnList, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleAPIResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// PartialUpdateServicesGgsnGtpSettings
//
// Operation ID: partialUpdateServicesGgsnGtpSettings
//
// Use this API command to modify the gtp setting of GGSN/PGW.
//
// Request Body:
//	 - body *WSGServiceGtpSettings
func (s *WSGGGSNPGWServiceService) PartialUpdateServicesGgsnGtpSettings(ctx context.Context, body *WSGServiceGtpSettings, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = apiRequestFromPool(http.MethodPatch, RouteWSGPartialUpdateServicesGgsnGtpSettings, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleAPIResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}
