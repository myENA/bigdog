package bigdog

// API Version: 1.0.0

import (
	"context"
	"encoding/json"
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

type SCIAlertSendNotification200ResponseType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIAlertSendNotification200ResponseType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SCIAlertSendNotification200ResponseType{XAdditionalProperties: tmp}
	return nil
}

func (t *SCIAlertSendNotification200ResponseType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSCIAlertSendNotification200ResponseType() *SCIAlertSendNotification200ResponseType {
	m := new(SCIAlertSendNotification200ResponseType)
	return m
}

// AlertSendNotification
func (s *SCIAlertService) AlertSendNotification(ctx context.Context, mutators ...RequestMutator) (*SCIAlertSendNotification200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIAlertSendNotification200ResponseType
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
	resp = NewSCIAlertSendNotification200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
