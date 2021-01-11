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
// Download CSV of resources within the VSZ
//
// Operation ID: exportCSVByQuery
// Operation path: /exportcsv/{resource}
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGCommonQueryCriteriaSuperSet
//
// Required parameters:
// - resource string
//		- required
//		- oneof:[alarm,event,ap,client,zone]
func (s *WSGCSVExportService) ExportCSVByQuery(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, resource string, mutators ...RequestMutator) (*FileAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newFileAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*FileAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGExportCSVByQuery, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*FileAPIResponse), err
	}
	req.PathParams.Set("resource", resource)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*FileAPIResponse), err
}
