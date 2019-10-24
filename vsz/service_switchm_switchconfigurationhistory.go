package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/switchm/common"
	"github.com/myENA/ruckus-client/vsz/types/switchm/deploylog"
	"github.com/myENA/ruckus-client/vsz/types/switchm/deploylogitem"
)

type SwitchMSwitchConfigurationHistoryService struct {
	client *Client
}

func NewSwitchMSwitchConfigurationHistoryService(client *Client) *SwitchMSwitchConfigurationHistoryService {
	s := new(SwitchMSwitchConfigurationHistoryService)
	s.client = client
	return s
}

func (ss *SwitchMService) SwitchMSwitchConfigurationHistoryService() *SwitchMSwitchConfigurationHistoryService {
	serv := new(SwitchMSwitchConfigurationHistoryService)
	serv.client = ss.client
	return serv
}

func (s *SwitchMSwitchConfigurationHistoryService) FindConfigurationHistory(ctx context.Context) (*deploylog.ConfigurationHistoryQueryResult, error) {
}

func (s *SwitchMSwitchConfigurationHistoryService) FindConfigurationHistoryByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) (*deploylog.ConfigurationHistoryQueryResult, error) {
}

func (s *SwitchMSwitchConfigurationHistoryService) FindConfigurationHistoryDetail(ctx context.Context) (*deploylogitem.ConfigurationHistoryDetailQueryResult, error) {
}

func (s *SwitchMSwitchConfigurationHistoryService) FindConfigurationHistoryDetailByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) (*deploylogitem.ConfigurationHistoryDetailQueryResult, error) {
}
