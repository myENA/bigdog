package bigdog

// API Version: v9_1

import (
	"context"
	"net/http"
)

type WSGCSVExportService struct {
	apiClient *VSZClient
}

func NewWSGCSVExportService(c *VSZClient) *WSGCSVExportService {
	s := new(WSGCSVExportService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGCSVExportService() *WSGCSVExportService {
	return NewWSGCSVExportService(ss.apiClient)
}

// ExportCSVByQuery
//
// Operation ID: exportCSVByQuery
//
// Download CSV of resources within the VSZ
//
// Request Body:
//	 - body string
//
// Required Parameters:
// - resource string
//		- required
//		- oneof:[alarm,event,ap,client,zone]
func (s *WSGCSVExportService) ExportCSVByQuery(ctx context.Context, body string, resource string, mutators ...RequestMutator) (*RawResponse, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *RawResponse
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGExportCSVByQuery, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, "application/octet-stream")
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetPathParameter("resource", resource)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = new(RawResponse)
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
