package ruckus

// API Version: 1.0.0

import (
	"context"
	"net/http"
)

type SCIPcireportService struct {
	apiClient *SCIClient
}

func NewSCIPcireportService(c *SCIClient) *SCIPcireportService {
	s := new(SCIPcireportService)
	s.apiClient = c
	return s
}

func (ss *SCIService) SCIPcireportService() *SCIPcireportService {
	return NewSCIPcireportService(ss.apiClient)
}

// PciReportDownloadReport
//
// Request Body:
//	 - body string
func (s *SCIPcireportService) PciReportDownloadReport(ctx context.Context, body string) (*APIResponseMeta, error) {
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
