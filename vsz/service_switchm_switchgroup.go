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

// AddGroup
//
// Use this API command to create a new switch group under an existing domain or switch group.
//
// Request Body:
//	 - body *group.SwitchGroup
func (s *SwitchMSwitchGroupService) AddGroup(ctx context.Context, body *group.SwitchGroup) (*group.AuditId, error) {
}

// FindGroupBySwitchGroupId
//
// Use this API command to retrieve switch group detail.
//
// Path Parameters:
// - pSwitchGroupId string
//		- required
func (s *SwitchMSwitchGroupService) FindGroupBySwitchGroupId(ctx context.Context, pSwitchGroupId string) (*group.SwitchGroupQueryResult, error) {
}

// FindGroupIdsByDomainByDomainId
//
// Use this API command to retrieve the switch groups by domain ID.
//
// Path Parameters:
// - pDomainId string
//		- required
func (s *SwitchMSwitchGroupService) FindGroupIdsByDomainByDomainId(ctx context.Context, pDomainId string) (*group.GroupsByIdsQueryResultList, error) {
}
