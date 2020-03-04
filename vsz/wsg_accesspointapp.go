package vsz

// API Version: v9_0

import (
	"context"
	"fmt"
	"net/http"
)

type WSGAccessPointAPPService struct {
	apiClient *APIClient
}

func NewWSGAccessPointAPPService(c *APIClient) *WSGAccessPointAPPService {
	s := new(WSGAccessPointAPPService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGAccessPointAPPService() *WSGAccessPointAPPService {
	return NewWSGAccessPointAPPService(ss.apiClient)
}

// FindApsLineman
//
// Use this API command to retrieve the summary information of an AP. This is used by the Ruckus Wireless AP mobile app.
//
// Optional Parameters:
// - domainId string
//		- nullable
// - index string
//		- nullable
// - listSize string
//		- nullable
// - showAlarm string
//		- nullable
// - zoneId string
//		- nullable
func (s *WSGAccessPointAPPService) FindApsLineman(ctx context.Context, optionalParams map[string][]string) (*WSGAPLinemanSummary, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGAPLinemanSummary
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, fmt.Sprintf("%s%s", s.apiClient.wsgPath, RouteWSGFindApsLineman), true)
	if v, ok := optionalParams["domainId"]; ok {
		req.AddQueryParameter("domainId", v)
	}
	if v, ok := optionalParams["index"]; ok {
		req.AddQueryParameter("index", v)
	}
	if v, ok := optionalParams["listSize"]; ok {
		req.AddQueryParameter("listSize", v)
	}
	if v, ok := optionalParams["showAlarm"]; ok {
		req.AddQueryParameter("showAlarm", v)
	}
	if v, ok := optionalParams["zoneId"]; ok {
		req.AddQueryParameter("zoneId", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGAPLinemanSummary()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// FindApsTotalCount
//
// Use this API command to retrieve the total AP count within a zone or a domain.
//
// Optional Parameters:
// - domainId string
//		- nullable
// - zoneId string
//		- nullable
func (s *WSGAccessPointAPPService) FindApsTotalCount(ctx context.Context, optionalParams map[string][]string) (interface{}, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     interface{}
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, fmt.Sprintf("%s%s", s.apiClient.wsgPath, RouteWSGFindApsTotalCount), true)
	if v, ok := optionalParams["domainId"]; ok {
		req.AddQueryParameter("domainId", v)
	}
	if v, ok := optionalParams["zoneId"]; ok {
		req.AddQueryParameter("zoneId", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = new(interface{})
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindLinemanWorkflow
//
// Use this API command to download the workflow file used by the Ruckus Wireless AP mobile app.
func (s *WSGAccessPointAPPService) FindLinemanWorkflow(ctx context.Context) ([]byte, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     []byte
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, fmt.Sprintf("%s%s", s.apiClient.wsgPath, RouteWSGFindLinemanWorkflow), true)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = make([]byte, 0)
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// UpdateLinemanWorkflow
//
// Use this API command to upload a workflow file used by the Ruckus Wireless AP mobile app.
//
// Request Body:
//	 - body []byte
func (s *WSGAccessPointAPPService) UpdateLinemanWorkflow(ctx context.Context, body []byte) (*APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodPut, fmt.Sprintf("%s%s", s.apiClient.wsgPath, RouteWSGUpdateLinemanWorkflow), true)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}
