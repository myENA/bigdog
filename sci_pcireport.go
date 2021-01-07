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
func (s *SCIPCIReportService) PciReportDownloadReport(ctx context.Context, formValues url.Values, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodPost, RouteSCIPciReportDownloadReport, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(bytes.NewBufferString(formValues.Encode())); err != nil {
		return resp.(*RawAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}
