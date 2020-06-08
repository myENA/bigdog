package ruckus

// API Version: v9_0

import (
	"context"
	"net/http"
)

type SwitchMJobandScheduleService struct {
	apiClient *VSZClient
}

func NewSwitchMJobandScheduleService(c *VSZClient) *SwitchMJobandScheduleService {
	s := new(SwitchMJobandScheduleService)
	s.apiClient = c
	return s
}

func (ss *SwitchMService) SwitchMJobandScheduleService() *SwitchMJobandScheduleService {
	return NewSwitchMJobandScheduleService(ss.apiClient)
}

// AddJob
//
// Use this API command to retrieve a list of jobs.
//
// Request Body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
func (s *SwitchMJobandScheduleService) AddJob(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet) (*SwitchMSwitchJobList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMSwitchJobList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSwitchMAddJob, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMSwitchJobList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// DeleteJobSchedule
//
// Use this API command to delete a selected schedule.
func (s *SwitchMJobandScheduleService) DeleteJobSchedule(ctx context.Context) (interface{}, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     interface{}
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteSwitchMDeleteJobSchedule, true)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = new(interface{})
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindJobByJobId
//
// Use this API command to retrieve a given job.
//
// Request Body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
//
// Required Parameters:
// - jobId string
//		- required
func (s *SwitchMJobandScheduleService) FindJobByJobId(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet, jobId string) (*SwitchMSwitchJobJob, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMSwitchJobJob
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteSwitchMFindJobByJobId, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetPathParameter("jobId", jobId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMSwitchJobJob()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindJobScheduleByScheduleId
//
// Use this API command to retrieve a given schedule.
//
// Required Parameters:
// - scheduleId string
//		- required
func (s *SwitchMJobandScheduleService) FindJobScheduleByScheduleId(ctx context.Context, scheduleId string) (*SwitchMSwitchJobJobScheduleResponse, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMSwitchJobJobScheduleResponse
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteSwitchMFindJobScheduleByScheduleId, true)
	req.SetPathParameter("scheduleId", scheduleId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMSwitchJobJobScheduleResponse()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
