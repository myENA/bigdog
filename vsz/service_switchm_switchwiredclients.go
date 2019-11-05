package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/switchm/common"
	"github.com/myENA/ruckus-client/vsz/types/switchm/switchmswitch"
)

type SwitchMSwitchWiredClientsService struct {
	apiClient *APIClient
}

func NewSwitchMSwitchWiredClientsService(c *APIClient) *SwitchMSwitchWiredClientsService {
	s := new(SwitchMSwitchWiredClientsService)
	s.apiClient = c
	return s
}

func (ss *SwitchMService) SwitchMSwitchWiredClientsService() *SwitchMSwitchWiredClientsService {
	serv := new(SwitchMSwitchWiredClientsService)
	serv.apiClient = ss.apiClient
	return serv
}

// AddSwitchClients
//
// Use this API command to retrieve all the wired clients connected to switch, currently managed by SmartZone.
//
// Request Body:
//	 - body *common.QueryCriteriaSuperSet
func (s *SwitchMSwitchWiredClientsService) AddSwitchClients(ctx context.Context, body *common.QueryCriteriaSuperSet) (*switchmswitch.ConnectedDevicesQueryList, error) {
}

// AddSwitchClientsAp
//
// Use this API command to retrieve all the Ruckus APs connected to switch, currently managed by SmartZone.
//
// Request Body:
//	 - body *common.QueryCriteriaSuperSet
func (s *SwitchMSwitchWiredClientsService) AddSwitchClientsAp(ctx context.Context, body *common.QueryCriteriaSuperSet) (*switchmswitch.ConnectedAPsQueryList, error) {
}
