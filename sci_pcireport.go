package bigdog

// API Version: 1.0.0

import (
	"bytes"
	"context"
	"net/http"
	"net/url"
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
//
// Form Data Parameters:
// - reportId string
//		- required
// - timezone string
//		- required
func (s *SCIPCIReportService) PciReportDownloadReport(ctx context.Context, formValues url.Values, mutators ...RequestMutator) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIPciReportDownloadReport, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, "*/*")
	if err = req.SetBody(bytes.NewBufferString(formValues.Encode())); err != nil {
		return rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}
