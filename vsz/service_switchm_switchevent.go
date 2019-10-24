package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/switchm/eventconfig"
)

type SwitchMSwitchEventService struct {
	client *Client
}

func NewSwitchMSwitchEventService(client *Client) *SwitchMSwitchEventService {
	s := new(SwitchMSwitchEventService)
	s.client = client
	return s
}

func (ss *SwitchMService) SwitchMSwitchEventService() *SwitchMSwitchEventService {
	serv := new(SwitchMSwitchEventService)
	serv.client = ss.client
	return serv
}

// FindCustomEvent
//
// Use this API command to retrieve switch event config list
func (s *SwitchMSwitchEventService) FindCustomEvent(ctx context.Context) (*eventconfig.GetEventConfigList, error) {
}

// FindCustomEventById
//
// Use this API command to retrieve one switch event config
func (s *SwitchMSwitchEventService) FindCustomEventById(ctx context.Context, pId string) (*eventconfig.EventConfig, error) {
}
