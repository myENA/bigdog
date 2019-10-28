package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/switchm/common"
	"github.com/myENA/ruckus-client/vsz/types/switchm/portsettings"
)

type SwitchMSwitchPortSettingsService struct {
	client *Client
}

func NewSwitchMSwitchPortSettingsService(client *Client) *SwitchMSwitchPortSettingsService {
	s := new(SwitchMSwitchPortSettingsService)
	s.client = client
	return s
}

func (ss *SwitchMService) SwitchMSwitchPortSettingsService() *SwitchMSwitchPortSettingsService {
	serv := new(SwitchMSwitchPortSettingsService)
	serv.client = ss.client
	return serv
}

// AddPortSettingsBulk
//
// Use this API command to Bulk update the port setting
//
// Request Body:
//	 - body *portsettings.CreateBulk
func (s *SwitchMSwitchPortSettingsService) AddPortSettingsBulk(ctx context.Context, body *portsettings.CreateBulk) (*portsettings.EmptyResult, error) {
}

// FindPortSettings
//
// Use this API command to Retrieve all Port Settings list.
func (s *SwitchMSwitchPortSettingsService) FindPortSettings(ctx context.Context) (*portsettings.PortSettingsQueryResult, error) {
}

// FindPortSettingsById
//
// Use this API command to Retrieve Port Settings.
//
// Path Parameters:
// - pId string
//		- required
func (s *SwitchMSwitchPortSettingsService) FindPortSettingsById(ctx context.Context, pId string) (*portsettings.PortSettings, error) {
}

// FindPortSettingsByQueryCriteria
//
// Use this API command to Retrieve Port Settings list.
//
// Request Body:
//	 - body *common.QueryCriteriaSuperSet
func (s *SwitchMSwitchPortSettingsService) FindPortSettingsByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) (*portsettings.PortSettingsQueryResult, error) {
}
