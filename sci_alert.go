package bigdog

// API Version: 1.0.0

import (
	"context"
	"net/http"
)

type SCIAlertService struct {
	apiClient *SCIClient
}

func NewSCIAlertService(c *SCIClient) *SCIAlertService {
	s := new(SCIAlertService)
	s.apiClient = c
	return s
}

func (ss *SCIService) SCIAlertService() *SCIAlertService {
	return NewSCIAlertService(ss.apiClient)
}

// AlertSendNotification
//
// Operation ID: alert.sendNotification
func (s *SCIAlertService) AlertSendNotification(ctx context.Context, mutators ...RequestMutator) (*RawResponse, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodPost, RouteSCIAlertSendNotification, true)
	req.SetHeader(headerKeyContentType, "*/*")
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = new(RawResponse)
	rm, err = handleRawResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
