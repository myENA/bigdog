package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/switchm/common"
	"github.com/myENA/ruckus-client/vsz/types/switchm/job"
)

type SwitchMJobandScheduleService struct {
	client *Client
}

func NewSwitchMJobandScheduleService(client *Client) *SwitchMJobandScheduleService {
	s := new(SwitchMJobandScheduleService)
	s.client = client
	return s
}

func (ss *SwitchMService) SwitchMJobandScheduleService() *SwitchMJobandScheduleService {
	serv := new(SwitchMJobandScheduleService)
	serv.client = ss.client
	return serv
}

// AddJob
//
// Use this API command to retrieve a list of jobs.
//
// Request Body:
//	 - body *common.QueryCriteriaSuperSet
func (s *SwitchMJobandScheduleService) AddJob(ctx context.Context, body *common.QueryCriteriaSuperSet) (*job.List, error) {
}

// DeleteJobSchedule
//
// Use this API command to delete a selected schedule.
func (s *SwitchMJobandScheduleService) DeleteJobSchedule(ctx context.Context) error {
}

// FindJobByJobId
//
// Use this API command to retrieve a given job.
//
// Request Body:
//	 - body *common.QueryCriteriaSuperSet
//
// Path Parameters:
// - pJobId string
//		- required
func (s *SwitchMJobandScheduleService) FindJobByJobId(ctx context.Context, body *common.QueryCriteriaSuperSet, pJobId string) (*job.Job, error) {
}

// FindJobScheduleByScheduleId
//
// Use this API command to retrieve a given schedule.
//
// Path Parameters:
// - pScheduleId string
//		- required
func (s *SwitchMJobandScheduleService) FindJobScheduleByScheduleId(ctx context.Context, pScheduleId string) (*job.JobScheduleResponse, error) {
}
