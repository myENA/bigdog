package vsz

// API Version: v8_1

import (
	"context"
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
func (s *WSGAccessPointAPPService) FindApsLineman(ctx context.Context, optionalParams map[string]interface{}) (*WSGAPLinemanSummary, error) {
	var (
		req  *APIRequest
		resp *WSGAPLinemanSummary
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindApsLineman, true)
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
func (s *WSGAccessPointAPPService) FindApsTotalCount(ctx context.Context, optionalParams map[string]interface{}) (interface{}, error) {
	var (
		req  *APIRequest
		resp interface{}
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindApsTotalCount, true)
	if v, ok := optionalParams["domainId"]; ok {
		req.AddQueryParameter("domainId", v)
	}
	if v, ok := optionalParams["zoneId"]; ok {
		req.AddQueryParameter("zoneId", v)
	}
}

// FindLinemanWorkflow
//
// Use this API command to download the workflow file used by the Ruckus Wireless AP mobile app.
func (s *WSGAccessPointAPPService) FindLinemanWorkflow(ctx context.Context) ([]byte, error) {
	var (
		req  *APIRequest
		resp []byte
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindLinemanWorkflow, true)
}

// UpdateLinemanWorkflow
//
// Use this API command to upload a workflow file used by the Ruckus Wireless AP mobile app.
//
// Request Body:
//	 - body []byte
func (s *WSGAccessPointAPPService) UpdateLinemanWorkflow(ctx context.Context, body []byte) error {
	var (
		req *APIRequest
		err error
	)
	if err = ctx.Err(); err != nil {
		return err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return err
	}
	req = NewAPIRequest(http.MethodPut, RouteWSGUpdateLinemanWorkflow, true)
	if err = req.SetBody(body); err != nil {
		return err
	}
}
