package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/switchm/common"
	"github.com/myENA/ruckus-client/vsz/types/switchm/switchmswitch"
)

type SwitchMSwitchService struct {
	client *Client
}

func NewSwitchMSwitchService(client *Client) *SwitchMSwitchService {
	s := new(SwitchMSwitchService)
	s.client = client
	return s
}

func (ss *SwitchMService) SwitchMSwitchService() *SwitchMSwitchService {
	serv := new(SwitchMSwitchService)
	serv.client = ss.client
	return serv
}

// AddSwitch
//
// Use this API command to retrieve all the switches currently managed by SmartZone.
//
// Request Body:
//	 - body *common.QueryCriteriaSuperSet
func (s *SwitchMSwitchService) AddSwitch(ctx context.Context, body *common.QueryCriteriaSuperSet) (*switchmswitch.SwitchQueryResultList, error) {
}

// AddSwitchSnmpSyncedSwitch
//
// Use this API command to retrieve all the switches currently managed by SmartZone and SNMP synced.
//
// Request Body:
//	 - body *common.QueryCriteriaSuperSet
func (s *SwitchMSwitchService) AddSwitchSnmpSyncedSwitch(ctx context.Context, body *common.QueryCriteriaSuperSet) (*switchmswitch.SwitchQueryResultList, error) {
}

// AddSwitchViewDetails
//
// Use this API command to retrieve switch and port details for the selected Switch/SwitchGroup/Domain.
//
// Request Body:
//	 - body *common.QueryCriteriaSuperSet
func (s *SwitchMSwitchService) AddSwitchViewDetails(ctx context.Context, body *common.QueryCriteriaSuperSet) (*switchmswitch.StackMemberQueryResult, error) {
}

// FindSwitchById
//
// Use this API command to retrieve a switch status.
//
// Path Parameters:
// - pId string
//		- required
func (s *SwitchMSwitchService) FindSwitchById(ctx context.Context, pId string) (*switchmswitch.NetworkSwitch, error) {
}

// FindSwitchFirmwareBySwitchId
//
// Use this API command to get a list of firmware update history.
//
// Path Parameters:
// - pSwitchId string
//		- required
func (s *SwitchMSwitchService) FindSwitchFirmwareBySwitchId(ctx context.Context, pSwitchId string) (*switchmswitch.FirmwareHistoryQueryResultList, error) {
}

// UpdateSwitchMoveByDestinationSwitchGroupId
//
// Use this API command to move a list of switches to a switch group.
//
// Request Body:
//	 - body switchmswitch.SwitchIdList
//
// Path Parameters:
// - pDestinationSwitchGroupId string
//		- required
func (s *SwitchMSwitchService) UpdateSwitchMoveByDestinationSwitchGroupId(ctx context.Context, body switchmswitch.SwitchIdList, pDestinationSwitchGroupId string) error {
}
