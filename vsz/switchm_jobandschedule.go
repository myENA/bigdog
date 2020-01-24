package vsz

// API Version: v8_1

import (
	"context"
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
func (s *SwitchMJobandScheduleService) AddJob(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet) (*SwitchMJobList, error) {
	var (
		resp *SwitchMJobList
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
	req := NewAPIRequest(http.MethodPost, RouteSwitchMAddJob, true)
}

// DeleteJobSchedule
//
// Use this API command to delete a selected schedule.
func (s *SwitchMJobandScheduleService) DeleteJobSchedule(ctx context.Context) (interface{}, error) {
	var (
		resp interface{}
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	req := NewAPIRequest(http.MethodDelete, RouteSwitchMDeleteJobSchedule, true)
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
func (s *SwitchMJobandScheduleService) FindJobByJobId(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet, jobId string) (*SwitchMJob, error) {
	var (
		resp *SwitchMJob
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, jobId, "required"); err != nil {
		return resp, err
	}
	req := NewAPIRequest(http.MethodGet, RouteSwitchMFindJobByJobId, true)
}

// FindJobScheduleByScheduleId
//
// Use this API command to retrieve a given schedule.
//
// Required Parameters:
// - scheduleId string
//		- required
func (s *SwitchMJobandScheduleService) FindJobScheduleByScheduleId(ctx context.Context, scheduleId string) (*SwitchMJobScheduleResponse, error) {
	var (
		resp *SwitchMJobScheduleResponse
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, scheduleId, "required"); err != nil {
		return resp, err
	}
	req := NewAPIRequest(http.MethodGet, RouteSwitchMFindJobScheduleByScheduleId, true)
}
