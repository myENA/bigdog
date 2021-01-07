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
//	 - body *WSGCommonQueryCriteriaSuperSet
//
// Required Parameters:
// - resource string
//		- required
//		- oneof:[alarm,event,ap,client,zone]
func (s *WSGCSVExportService) ExportCSVByQuery(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, resource string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGExportCSVByQuery, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req.PathParams.Set("resource", resource)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}
