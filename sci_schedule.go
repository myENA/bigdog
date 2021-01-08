package bigdog

// API Version: 1.0.0

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
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
// Definition: schedule_batchDelete200ResponseType
type SCIScheduleBatchDelete200ResponseType struct {
	Count *float64 `json:"count,omitempty"`
}

type SCIScheduleBatchDelete200ResponseTypeAPIResponse struct {
	*RawAPIResponse
	Data *SCIScheduleBatchDelete200ResponseType
}

func newSCIScheduleBatchDelete200ResponseTypeAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIScheduleBatchDelete200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIScheduleBatchDelete200ResponseTypeAPIResponse) Hydrate() error {
	r.Data = new(SCIScheduleBatchDelete200ResponseType)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSCIScheduleBatchDelete200ResponseType() *SCIScheduleBatchDelete200ResponseType {
	m := new(SCIScheduleBatchDelete200ResponseType)
	return m
}

// ScheduleBatchDelete
//
// Operation ID: schedule_batchDelete
//
// Delete schedules with their related filters and occurrences in a single transaction.
//
// Form Data Parameters:
// - ids string
//		- required
func (s *SCIScheduleService) ScheduleBatchDelete(ctx context.Context, formValues url.Values, mutators ...RequestMutator) (*SCIScheduleBatchDelete200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSCIScheduleBatchDelete200ResponseTypeAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SCIScheduleBatchDelete200ResponseTypeAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteSCIScheduleBatchDelete, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(bytes.NewBufferString(formValues.Encode())); err != nil {
		return resp.(*SCIScheduleBatchDelete200ResponseTypeAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIScheduleBatchDelete200ResponseTypeAPIResponse), err
}

// ScheduleCreateWithRelations
//
// Operation ID: schedule_createWithRelations
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
func (s *SCIScheduleService) ScheduleCreateWithRelations(ctx context.Context, formValues url.Values, mutators ...RequestMutator) (*SCIModelsScheduleAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSCIModelsScheduleAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SCIModelsScheduleAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteSCIScheduleCreateWithRelations, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(bytes.NewBufferString(formValues.Encode())); err != nil {
		return resp.(*SCIModelsScheduleAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIModelsScheduleAPIResponse), err
}

// ScheduleExecuteJob
//
// Operation ID: schedule_executeJob
//
// Run schedule job
func (s *SCIScheduleService) ScheduleExecuteJob(ctx context.Context, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodPost, RouteSCIScheduleExecuteJob, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// ScheduleUpdateWithRelations
//
// Operation ID: schedule_updateWithRelations
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
func (s *SCIScheduleService) ScheduleUpdateWithRelations(ctx context.Context, formValues url.Values, id string, mutators ...RequestMutator) (*SCIModelsScheduleAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSCIModelsScheduleAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SCIModelsScheduleAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPut, RouteSCIScheduleUpdateWithRelations, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(bytes.NewBufferString(formValues.Encode())); err != nil {
		return resp.(*SCIModelsScheduleAPIResponse), err
	}
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIModelsScheduleAPIResponse), err
}
