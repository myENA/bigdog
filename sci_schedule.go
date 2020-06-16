package bigdog

// API Version: 1.0.0

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/url"
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

// SCIScheduleBatchDelete200ResponseType
//
// Definition: schedule.batchDelete200ResponseType
type SCIScheduleBatchDelete200ResponseType struct {
	Count *float64 `json:"count,omitempty"`
}

func NewSCIScheduleBatchDelete200ResponseType() *SCIScheduleBatchDelete200ResponseType {
	m := new(SCIScheduleBatchDelete200ResponseType)
	return m
}

// SCIScheduleExecuteJob200ResponseType
//
// Definition: schedule.executeJob200ResponseType
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
// Operation ID: schedule.batchDelete
//
// Delete schedules with their related filters and occurrences in a single transaction.
//
// Form Data Parameters:
// - ids string
//		- required
func (s *SCIScheduleService) ScheduleBatchDelete(ctx context.Context, formValues url.Values, mutators ...RequestMutator) (*SCIScheduleBatchDelete200ResponseType, *APIResponseMeta, error) {
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
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(bytes.NewBufferString(formValues.Encode())); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIScheduleBatchDelete200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ScheduleCreateWithRelations
//
// Operation ID: schedule.createWithRelations
//
// Create schedule with filter and occurrence in a single transaction.
//
// Form Data Parameters:
// - filterData string
//		- required
// - reportId string
//		- required
// - scheduleData string
//		- required
func (s *SCIScheduleService) ScheduleCreateWithRelations(ctx context.Context, formValues url.Values, mutators ...RequestMutator) (*SCIModelsSchedule, *APIResponseMeta, error) {
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
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(bytes.NewBufferString(formValues.Encode())); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIModelsSchedule()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ScheduleExecuteJob
//
// Operation ID: schedule.executeJob
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
	req.SetHeader(headerKeyContentType, "*/*")
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIScheduleExecuteJob200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ScheduleUpdateWithRelations
//
// Operation ID: schedule.updateWithRelations
//
// Update schedule with filter and occurrence in a single transaction.
//
// Form Data Parameters:
// - filterData string
//		- required
// - scheduleData string
//		- required
//
// Required Parameters:
// - id string
//		- required
func (s *SCIScheduleService) ScheduleUpdateWithRelations(ctx context.Context, formValues url.Values, id string, mutators ...RequestMutator) (*SCIModelsSchedule, *APIResponseMeta, error) {
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
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(bytes.NewBufferString(formValues.Encode())); err != nil {
		return resp, rm, err
	}
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIModelsSchedule()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
