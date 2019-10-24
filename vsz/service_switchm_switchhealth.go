package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/switchm/common"
	"github.com/myENA/ruckus-client/vsz/types/switchm/health"
)

type SwitchMSwitchHealthService struct {
	client *Client
}

func NewSwitchMSwitchHealthService(client *Client) *SwitchMSwitchHealthService {
	s := new(SwitchMSwitchHealthService)
	s.client = client
	return s
}

func (ss *SwitchMService) SwitchMSwitchHealthService() *SwitchMSwitchHealthService {
	serv := new(SwitchMSwitchHealthService)
	serv.client = ss.client
	return serv
}

// AddHealthCpuAgg
//
// Use this API command to retrieve aggregated CPU (min, max, avg, curr) data based on the time duration.
//
// Request Body:
//	 - body *common.QueryCriteriaSuperSet
func (s *SwitchMSwitchHealthService) AddHealthCpuAgg(ctx context.Context, body *common.QueryCriteriaSuperSet) (*health.AggMetrics, error) {
}

// AddHealthCpuLine
//
// Use this API command to retrieve CPU trend data based on the time duration.
//
// Request Body:
//	 - body *common.QueryCriteriaSuperSet
func (s *SwitchMSwitchHealthService) AddHealthCpuLine(ctx context.Context, body *common.QueryCriteriaSuperSet) (*health.IcxMetrics, error) {
}

// AddHealthMemAgg
//
// Use this API command to retrieve aggregated CPU (min, max, avg, curr) data based on the time duration.
//
// Request Body:
//	 - body *common.QueryCriteriaSuperSet
func (s *SwitchMSwitchHealthService) AddHealthMemAgg(ctx context.Context, body *common.QueryCriteriaSuperSet) (*health.AggMetrics, error) {
}

// AddHealthMemLine
//
// Use this API command to retrieve switch memory trend data based on the time duration.
//
// Request Body:
//	 - body *common.QueryCriteriaSuperSet
func (s *SwitchMSwitchHealthService) AddHealthMemLine(ctx context.Context, body *common.QueryCriteriaSuperSet) (*health.IcxMetrics, error) {
}

// AddHealthStatus
//
// Use this API command to retrieve switch health status.
//
// Request Body:
//	 - body *common.QueryCriteriaSuperSet
func (s *SwitchMSwitchHealthService) AddHealthStatus(ctx context.Context, body *common.QueryCriteriaSuperSet) (*health.Status, error) {
}

// AddHealthStatusAll
//
// Use this API command to retrieve fan, temperature and power supply status for the switch managed by SmartZone.
//
// Request Body:
//	 - body *common.QueryCriteriaSuperSet
func (s *SwitchMSwitchHealthService) AddHealthStatusAll(ctx context.Context, body *common.QueryCriteriaSuperSet) (*health.Status, error) {
}
