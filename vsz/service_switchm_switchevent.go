package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/switchm/eventconfig"
)

type SwitchMSwitchEventService struct {
	apiClient *APIClient
}

func NewSwitchMSwitchEventService(c *APIClient) *SwitchMSwitchEventService {
	s := new(SwitchMSwitchEventService)
	s.apiClient = c
	return s
}

func (ss *SwitchMService) SwitchMSwitchEventService() *SwitchMSwitchEventService {
	serv := new(SwitchMSwitchEventService)
	serv.apiClient = ss.apiClient
	return serv
}

// AddCustomEvent
//
// Use this API command to create a new text pattern event config
//
// Request Body:
//	 - body *eventconfig.AddEventConfig
func (s *SwitchMSwitchEventService) AddCustomEvent(ctx context.Context, body *eventconfig.AddEventConfig) (*eventconfig.AddEventConfigResult, error) {
}

// DeleteCustomEventById
//
// Use this API command to delete a text pattern event config
//
// Path Parameters:
// - pId string
//		- required
func (s *SwitchMSwitchEventService) DeleteCustomEventById(ctx context.Context, pId string) (*eventconfig.DeleteEventConfigResult, error) {
}

// FindCustomEvent
//
// Use this API command to retrieve switch event config list
func (s *SwitchMSwitchEventService) FindCustomEvent(ctx context.Context) (*eventconfig.GetEventConfigList, error) {
}

// FindCustomEventById
//
// Use this API command to retrieve one switch event config
//
// Path Parameters:
// - pId string
//		- required
func (s *SwitchMSwitchEventService) FindCustomEventById(ctx context.Context, pId string) (*eventconfig.EventConfig, error) {
}

// UpdateCustomEventById
//
// Use this API command to modify a switch custom event config. The patch variable {id} is same as id attribute in the request payload. For CPU/Memory, only key, type, criteria, and severity attributes are required.
//
// Request Body:
//	 - body *eventconfig.UpdateEventConfig
//
// Path Parameters:
// - pId string
//		- required
func (s *SwitchMSwitchEventService) UpdateCustomEventById(ctx context.Context, body *eventconfig.UpdateEventConfig, pId string) (*eventconfig.UpdateEventConfigResult, error) {
}
