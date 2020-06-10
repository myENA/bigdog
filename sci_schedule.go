package bigdog

// API Version: 1.0.0

import (
	"context"
	"encoding/json"
	"net/http"
)

type SCIScheduleService struct {
	apiClient *SCIClient
}

func NewSCIScheduleService(c *SCIClient) *SCIScheduleService {
	s := new(SCIScheduleService)
	s.apiClient = c
	return s
}

func (ss *SCIService) SCIScheduleService() *SCIScheduleService {
	return NewSCIScheduleService(ss.apiClient)
}

type SCIScheduleBatchDelete200ResponseType struct {
	Count *float64 `json:"count,omitempty"`
}

func NewSCIScheduleBatchDelete200ResponseType() *SCIScheduleBatchDelete200ResponseType {
	m := new(SCIScheduleBatchDelete200ResponseType)
	return m
}

type SCIScheduleExecuteJob200ResponseType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIScheduleExecuteJob200ResponseType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SCIScheduleExecuteJob200ResponseType{XAdditionalProperties: tmp}
	return nil
}

func (t *SCIScheduleExecuteJob200ResponseType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSCIScheduleExecuteJob200ResponseType() *SCIScheduleExecuteJob200ResponseType {
	m := new(SCIScheduleExecuteJob200ResponseType)
	return m
}

// ScheduleBatchDelete
//
// Delete schedules with their related filters and occurrences in a single transaction.
//
// Request Body:
//	 - body string
func (s *SCIScheduleService) ScheduleBatchDelete(ctx context.Context, ids string, mutators ...RequestMutator) (*SCIScheduleBatchDelete200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIScheduleBatchDelete200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIScheduleBatchDelete, true)
	if err = req.SetBody(ids); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIScheduleBatchDelete200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ScheduleCreateWithRelations
//
// Create schedule with filter and occurrence in a single transaction.
//
// Request Body:
//	 - body string
func (s *SCIScheduleService) ScheduleCreateWithRelations(ctx context.Context, reportId string, mutators ...RequestMutator) (*SCIModelsSchedule, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIModelsSchedule
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIScheduleCreateWithRelations, true)
	if err = req.SetBody(reportId); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIModelsSchedule()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ScheduleExecuteJob
//
// Run schedule job
func (s *SCIScheduleService) ScheduleExecuteJob(ctx context.Context, mutators ...RequestMutator) (*SCIScheduleExecuteJob200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIScheduleExecuteJob200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIScheduleExecuteJob, true)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIScheduleExecuteJob200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ScheduleUpdateWithRelations
//
// Update schedule with filter and occurrence in a single transaction.
//
// Request Body:
//	 - body string
//
// Required Parameters:
// - id string
//		- required
func (s *SCIScheduleService) ScheduleUpdateWithRelations(ctx context.Context, scheduleData string, id string, mutators ...RequestMutator) (*SCIModelsSchedule, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIModelsSchedule
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPut, RouteSCIScheduleUpdateWithRelations, true)
	if err = req.SetBody(scheduleData); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIModelsSchedule()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
