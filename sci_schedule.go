package bigdog

// API Version: 1.0.0

import (
	"bytes"
	"context"
	"errors"
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

func newSCIScheduleBatchDelete200ResponseTypeAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SCIScheduleBatchDelete200ResponseTypeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *SCIScheduleBatchDelete200ResponseTypeAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(SCIScheduleBatchDelete200ResponseType)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewSCIScheduleBatchDelete200ResponseType() *SCIScheduleBatchDelete200ResponseType {
	m := new(SCIScheduleBatchDelete200ResponseType)
	return m
}

// ScheduleBatchDelete
//
// Delete schedules with their related filters and occurrences in a single transaction.
//
// Operation ID: schedule_batchDelete
// Operation path: /schedules/batchDelete
// Success code: 200 (OK)
//
// Form data parameters:
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
	req = apiRequestFromPool(APISourceSCI, http.MethodPost, RouteSCIScheduleBatchDelete, true)
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
// Create schedule with filter and occurrence in a single transaction.
//
// Operation ID: schedule_createWithRelations
// Operation path: /schedules/createWithRelations
// Success code: 200 (OK)
//
// Form data parameters:
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
	req = apiRequestFromPool(APISourceSCI, http.MethodPost, RouteSCIScheduleCreateWithRelations, true)
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
// Run schedule job
//
// Operation ID: schedule_executeJob
// Operation path: /schedules/executeJob
// Success code: 200 (OK)
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
	req = apiRequestFromPool(APISourceSCI, http.MethodPost, RouteSCIScheduleExecuteJob, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// ScheduleUpdateWithRelations
//
// Update schedule with filter and occurrence in a single transaction.
//
// Operation ID: schedule_updateWithRelations
// Operation path: /schedules/{id}/updateWithRelations
// Success code: 200 (OK)
//
// Form data parameters:
// - filterData string
//		- required
// - scheduleData string
//		- required
//
// Required parameters:
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
	req = apiRequestFromPool(APISourceSCI, http.MethodPut, RouteSCIScheduleUpdateWithRelations, true)
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
