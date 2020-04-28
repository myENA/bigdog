package vsz

// API Version: v9_0

import (
	"context"
	"net/http"
)

type WSGGGSNPGWServiceService struct {
	apiClient *APIClient
}

func NewWSGGGSNPGWServiceService(c *APIClient) *WSGGGSNPGWServiceService {
	s := new(WSGGGSNPGWServiceService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGGGSNPGWServiceService() *WSGGGSNPGWServiceService {
	return NewWSGGGSNPGWServiceService(ss.apiClient)
}

// DeleteServicesGgsnDnsServerList
//
// Use this API command to Disable the dns server list of GGSN/PGW.
func (s *WSGGGSNPGWServiceService) DeleteServicesGgsnDnsServerList(ctx context.Context) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteServicesGgsnDnsServerList, true)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteServicesGgsnGgsnList
//
// Use this API command to disable the ggsn server list of GGSN/PGW.
func (s *WSGGGSNPGWServiceService) DeleteServicesGgsnGgsnList(ctx context.Context) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteServicesGgsnGgsnList, true)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// FindGgsnGtpcConStatsByQueryCriteria
//
// Use this API command to retrieve a list of GGSN Connection.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGGGSNPGWServiceService) FindGgsnGtpcConStatsByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet) (*WSGRacStatsGgsnGtpcConList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGRacStatsGgsnGtpcConList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, rm, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGFindGgsnGtpcConStatsByQueryCriteria, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGRacStatsGgsnGtpcConList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindGgsnGtpStatsByQueryCriteria
//
// Use this API command to retrieve a list of GGSN/PGW GTP-C Sessions.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGGGSNPGWServiceService) FindGgsnGtpStatsByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet) (*WSGRacStatsGgsnGtpList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGRacStatsGgsnGtpList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, rm, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGFindGgsnGtpStatsByQueryCriteria, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGRacStatsGgsnGtpList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindServicesGgsn
//
// Use this API command to retrieve GGSN/PGW setting.
func (s *WSGGGSNPGWServiceService) FindServicesGgsn(ctx context.Context) (*WSGServiceGgsnConfig, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodGet, RouteWSGFindServicesGgsn, true)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGServiceGgsnConfig()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// PartialUpdateServicesGgsn
//
// Use this API command to modify the setting of GGSN/PGW.
//
// Request Body:
//	 - body *WSGServiceGgsnConfig
func (s *WSGGGSNPGWServiceService) PartialUpdateServicesGgsn(ctx context.Context, body *WSGServiceGgsnConfig) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return rm, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateServicesGgsn, true)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// PartialUpdateServicesGgsnDnsServerList
//
// Use this API command to modify the dns server list of GGSN/PGW.
//
// Request Body:
//	 - body WSGServiceDnsServerList
func (s *WSGGGSNPGWServiceService) PartialUpdateServicesGgsnDnsServerList(ctx context.Context, body WSGServiceDnsServerList) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateServicesGgsnDnsServerList, true)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// PartialUpdateServicesGgsnGgsnList
//
// Use this API command to modify the ggsn server list of GGSN/PGW.
//
// Request Body:
//	 - body WSGServiceGgsnList
func (s *WSGGGSNPGWServiceService) PartialUpdateServicesGgsnGgsnList(ctx context.Context, body WSGServiceGgsnList) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateServicesGgsnGgsnList, true)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// PartialUpdateServicesGgsnGtpSettings
//
// Use this API command to modify the gtp setting of GGSN/PGW.
//
// Request Body:
//	 - body *WSGServiceGtpSettings
func (s *WSGGGSNPGWServiceService) PartialUpdateServicesGgsnGtpSettings(ctx context.Context, body *WSGServiceGtpSettings) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return rm, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateServicesGgsnGtpSettings, true)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}
