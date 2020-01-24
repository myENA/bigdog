package vsz

// API Version: v8_1

import (
	"context"
	"net/http"
)

type SwitchMEventService struct {
	apiClient *APIClient
}

func NewSwitchMEventService(c *APIClient) *SwitchMEventService {
	s := new(SwitchMEventService)
	s.apiClient = c
	return s
}

func (ss *SwitchMService) SwitchMEventService() *SwitchMEventService {
	return NewSwitchMEventService(ss.apiClient)
}

// AddCustomEvent
//
// Use this API command to create a new text pattern event config
//
// Request Body:
//	 - body *SwitchMEventConfigAddEventConfig
func (s *SwitchMEventService) AddCustomEvent(ctx context.Context, body *SwitchMEventConfigAddEventConfig) (*SwitchMEventConfigAddEventConfigResult, error) {
	var (
		resp *SwitchMEventConfigAddEventConfigResult
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	req := NewAPIRequest(http.MethodPost, RouteSwitchMAddCustomEvent, true)
}

// DeleteCustomEventById
//
// Use this API command to delete a text pattern event config
//
// Required Parameters:
// - id string
//		- required
func (s *SwitchMEventService) DeleteCustomEventById(ctx context.Context, id string) (*SwitchMEventConfigDeleteEventConfigResult, error) {
	var (
		resp *SwitchMEventConfigDeleteEventConfigResult
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req := NewAPIRequest(http.MethodDelete, RouteSwitchMDeleteCustomEventById, true)
}

// FindCustomEvent
//
// Use this API command to retrieve switch event config list
func (s *SwitchMEventService) FindCustomEvent(ctx context.Context) (*SwitchMEventConfigGetEventConfigList, error) {
	var (
		resp *SwitchMEventConfigGetEventConfigList
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	req := NewAPIRequest(http.MethodGet, RouteSwitchMFindCustomEvent, true)
}

// FindCustomEventById
//
// Use this API command to retrieve one switch event config
//
// Required Parameters:
// - id string
//		- required
func (s *SwitchMEventService) FindCustomEventById(ctx context.Context, id string) (*SwitchMEventConfig, error) {
	var (
		resp *SwitchMEventConfig
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req := NewAPIRequest(http.MethodGet, RouteSwitchMFindCustomEventById, true)
}

// UpdateCustomEventById
//
// Use this API command to modify a switch custom event config. The patch variable {id} is same as id attribute in the request payload. For CPU/Memory, only key, type, criteria, and severity attributes are required.
//
// Request Body:
//	 - body *SwitchMEventConfigUpdateEventConfig
//
// Required Parameters:
// - id string
//		- required
func (s *SwitchMEventService) UpdateCustomEventById(ctx context.Context, body *SwitchMEventConfigUpdateEventConfig, id string) (*SwitchMEventConfigUpdateEventConfigResult, error) {
	var (
		resp *SwitchMEventConfigUpdateEventConfigResult
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req := NewAPIRequest(http.MethodPut, RouteSwitchMUpdateCustomEventById, true)
}
