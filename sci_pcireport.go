package bigdog

// API Version: 1.0.0

import (
	"context"
	"net/http"
	"net/url"
	"time"
)

type SCIPCIReportService struct {
	apiClient *SCIClient
}

func NewSCIPCIReportService(c *SCIClient) *SCIPCIReportService {
	s := new(SCIPCIReportService)
	s.apiClient = c
	return s
}

func (ss *SCIService) SCIPCIReportService() *SCIPCIReportService {
	return NewSCIPCIReportService(ss.apiClient)
}

// PciReportDownloadReport
//
// Operation ID: pciReport_downloadReport
// Operation path: /pciReports/download
// Success code: 204 (No Content)
//
// Form data parameters:
// - reportId string
//		- required
// - timezone string
//		- required
func (s *SCIPCIReportService) PciReportDownloadReport(ctx context.Context, formValues url.Values, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceSCI, http.MethodPost, RouteSCIPciReportDownloadReport, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(formValues)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*EmptyAPIResponse), err
}
