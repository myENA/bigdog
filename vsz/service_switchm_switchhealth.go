package vsz

// API Version: v8_0

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/switchm/common"
	"github.com/myENA/ruckus-client/vsz/types/switchm/health"
)

type SwitchMSwitchHealthService struct {
    client *Client
}

func NewSwitchMSwitchHealthService (client *Client) *SwitchMSwitchHealthService {
    s := new(SwitchMSwitchHealthService)
    s.client = client
    return s
}

func (ss *SwitchMService) SwitchMSwitchHealthService () *SwitchMSwitchHealthService {
    serv := new(SwitchMSwitchHealthService)
    serv.client = ss.client
    return serv
}

func (s *SwitchMSwitchHealthService) AddHealthCpuAgg (ctx context.Context) (health.AggMetrics, error) {
}

func (s *SwitchMSwitchHealthService) AddHealthCpuLine (ctx context.Context) (health.IcxMetrics, error) {
}

func (s *SwitchMSwitchHealthService) AddHealthMemAgg (ctx context.Context) (health.AggMetrics, error) {
}

func (s *SwitchMSwitchHealthService) AddHealthMemLine (ctx context.Context) (health.IcxMetrics, error) {
}

func (s *SwitchMSwitchHealthService) AddHealthStatus (ctx context.Context) (health.Status, error) {
}

func (s *SwitchMSwitchHealthService) AddHealthStatusAll (ctx context.Context) (health.Status, error) {
}

