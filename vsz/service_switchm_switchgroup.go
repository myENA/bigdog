package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/switchm/group"
)

type SwitchMSwitchGroupService struct {
	client *Client
}

func NewSwitchMSwitchGroupService(client *Client) *SwitchMSwitchGroupService {
	s := new(SwitchMSwitchGroupService)
	s.client = client
	return s
}

func (ss *SwitchMService) SwitchMSwitchGroupService() *SwitchMSwitchGroupService {
	serv := new(SwitchMSwitchGroupService)
	serv.client = ss.client
	return serv
}

func (s *SwitchMSwitchGroupService) AddGroup(ctx context.Context, body *group.SwitchGroup) (*group.AuditId, error) {
}

func (s *SwitchMSwitchGroupService) FindGroupBySwitchGroupId(ctx context.Context, pSwitchGroupId string) (*group.SwitchGroupQueryResult, error) {
}

func (s *SwitchMSwitchGroupService) FindGroupIdsByDomainByDomainId(ctx context.Context, pDomainId string) (*group.GroupsByIdsQueryResultList, error) {
}
