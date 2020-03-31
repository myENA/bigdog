package vsz

// API Version: v9_0

import (
	"context"
	"net/http"
)

type WSGSyslogServerService struct {
	apiClient *APIClient
}

func NewWSGSyslogServerService(c *APIClient) *WSGSyslogServerService {
	s := new(WSGSyslogServerService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGSyslogServerService() *WSGSyslogServerService {
	return NewWSGSyslogServerService(ss.apiClient)
}

// FindSystemSyslog
//
// Retrieve syslog server sertting.
//
// Optional Parameters:
// - index string
//		- nullable
// - listSize string
//		- nullable
func (s *WSGSyslogServerService) FindSystemSyslog(ctx context.Context, optionalParams map[string][]string) (*WSGSyslogServerSetting, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGSyslogServerSetting
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindSystemSyslog, true)
	if v, ok := optionalParams["index"]; ok {
		req.AddQueryParameter("index", v)
	}
	if v, ok := optionalParams["listSize"]; ok {
		req.AddQueryParameter("listSize", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGSyslogServerSetting()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// PartialUpdateSystemSyslog
//
// Modify syslog server setting.
//
// Request Body:
//	 - body *WSGSyslogModifySyslogSettings
func (s *WSGSyslogServerService) PartialUpdateSystemSyslog(ctx context.Context, body *WSGSyslogModifySyslogSettings) (*APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateSystemSyslog, true)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// PartialUpdateSystemSyslogPrimaryServer
//
// Modify Primary Server of syslog.
//
// Request Body:
//	 - body *WSGSyslogPrimaryServer
func (s *WSGSyslogServerService) PartialUpdateSystemSyslogPrimaryServer(ctx context.Context, body *WSGSyslogPrimaryServer) (*APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateSystemSyslogPrimaryServer, true)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// PartialUpdateSystemSyslogPriority
//
// Modify Priority of syslog.
//
// Request Body:
//	 - body *WSGSyslogPriority
func (s *WSGSyslogServerService) PartialUpdateSystemSyslogPriority(ctx context.Context, body *WSGSyslogPriority) (*APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateSystemSyslogPriority, true)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// PartialUpdateSystemSyslogSecondaryServer
//
// Modify Secondary Server of syslog.
//
// Request Body:
//	 - body *WSGSyslogSecondaryServer
func (s *WSGSyslogServerService) PartialUpdateSystemSyslogSecondaryServer(ctx context.Context, body *WSGSyslogSecondaryServer) (*APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateSystemSyslogSecondaryServer, true)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}
