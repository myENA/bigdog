package vsz

// API Version: v9_0

import (
	"context"
	"fmt"
	"net/http"
)

type SwitchMJobandScheduleService struct {
	apiClient *APIClient
}

func NewSwitchMJobandScheduleService(c *APIClient) *SwitchMJobandScheduleService {
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
func (s *SwitchMJobandScheduleService) AddJob(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet) (*SwitchMJobList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMJobList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, rm, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, fmt.Sprintf("%s%s", s.apiClient.switchMPath, RouteSwitchMAddJob), true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMJobList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
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
	req = NewAPIRequest(http.MethodDelete, fmt.Sprintf("%s%s", s.apiClient.switchMPath, RouteSwitchMDeleteJobSchedule), true)
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
func (s *SwitchMJobandScheduleService) FindJobByJobId(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet, jobId string) (*SwitchMJob, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMJob
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, rm, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, jobId, "required"); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, fmt.Sprintf("%s%s", s.apiClient.switchMPath, RouteSwitchMFindJobByJobId), true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetPathParameter("jobId", jobId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMJob()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// FindJobScheduleByScheduleId
//
// Use this API command to retrieve a given schedule.
//
// Required Parameters:
// - scheduleId string
//		- required
func (s *SwitchMJobandScheduleService) FindJobScheduleByScheduleId(ctx context.Context, scheduleId string) (*SwitchMJobScheduleResponse, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMJobScheduleResponse
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, scheduleId, "required"); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, fmt.Sprintf("%s%s", s.apiClient.switchMPath, RouteSwitchMFindJobScheduleByScheduleId), true)
	req.SetPathParameter("scheduleId", scheduleId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMJobScheduleResponse()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}
