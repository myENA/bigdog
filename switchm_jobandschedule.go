package bigdog

// API Version: v9_1

import (
	"context"
	"net/http"
)

type SwitchMJobAndScheduleService struct {
	apiClient *VSZClient
}

func NewSwitchMJobAndScheduleService(c *VSZClient) *SwitchMJobAndScheduleService {
	s := new(SwitchMJobAndScheduleService)
	s.apiClient = c
	return s
}

func (ss *SwitchMService) SwitchMJobAndScheduleService() *SwitchMJobAndScheduleService {
	return NewSwitchMJobAndScheduleService(ss.apiClient)
}

// AddJob
//
// Use this API command to retrieve a list of jobs.
//
// Operation ID: addJob
// Operation path: /job
// Success code: 200 (OK)
//
// Request body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
func (s *SwitchMJobAndScheduleService) AddJob(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*SwitchMJobListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSwitchMJobListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SwitchMJobListAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteSwitchMAddJob, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*SwitchMJobListAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SwitchMJobListAPIResponse), err
}

// DeleteJobSchedule
//
// Use this API command to delete a selected schedule.
//
// Operation ID: deleteJobSchedule
// Operation path: /job/schedule
// Success code: 200 (OK)
func (s *SwitchMJobAndScheduleService) DeleteJobSchedule(ctx context.Context, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodDelete, RouteSwitchMDeleteJobSchedule, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// FindJobByJobId
//
// Use this API command to retrieve a given job.
//
// Operation ID: findJobByJobId
// Operation path: /job/{jobId}
// Success code: 200 (OK)
//
// Request body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
//
// Required parameters:
// - jobId string
//		- required
func (s *SwitchMJobAndScheduleService) FindJobByJobId(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet, jobId string, mutators ...RequestMutator) (*SwitchMJobAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSwitchMJobAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SwitchMJobAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteSwitchMFindJobByJobId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*SwitchMJobAPIResponse), err
	}
	req.PathParams.Set("jobId", jobId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SwitchMJobAPIResponse), err
}

// FindJobScheduleByScheduleId
//
// Use this API command to retrieve a given schedule.
//
// Operation ID: findJobScheduleByScheduleId
// Operation path: /job/schedule/{scheduleId}
// Success code: 200 (OK)
//
// Required parameters:
// - scheduleId string
//		- required
func (s *SwitchMJobAndScheduleService) FindJobScheduleByScheduleId(ctx context.Context, scheduleId string, mutators ...RequestMutator) (*SwitchMJobScheduleResponseAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSwitchMJobScheduleResponseAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SwitchMJobScheduleResponseAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteSwitchMFindJobScheduleByScheduleId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("scheduleId", scheduleId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SwitchMJobScheduleResponseAPIResponse), err
}
