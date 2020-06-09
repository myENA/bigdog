package bigdog

// API Version: 1.0.0

import (
	"context"
	"net/http"
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
// Request Body:
//	 - body string
func (s *SCIPCIReportService) PciReportDownloadReport(ctx context.Context, body string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIPciReportDownloadReport, false)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}
