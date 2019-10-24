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

func (s *SwitchMJobandScheduleService) AddJob(ctx context.Context, body *common.QueryCriteriaSuperSet) (*job.List, error) {
}

func (s *SwitchMJobandScheduleService) DeleteJobSchedule(ctx context.Context) error {
}

func (s *SwitchMJobandScheduleService) FindJobByJobId(ctx context.Context, body *common.QueryCriteriaSuperSet, pJobId string) (*job.Job, error) {
}

func (s *SwitchMJobandScheduleService) FindJobScheduleByScheduleId(ctx context.Context, pScheduleId string) (*job.JobScheduleResponse, error) {
}
