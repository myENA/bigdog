package bigdog

// API Version: v9_1

import (
	"context"
	"net/http"
	"time"
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
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newFileAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGExportCSVByQuery, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	req.PathParams.Set("resource", resource)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*FileAPIResponse), err
}
