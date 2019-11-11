package vsz

// API Version: v8_1

import (
	"context"
	"errors"
	"fmt"
	"gopkg.in/go-playground/validator.v9"
)

type SwitchMTopReportsService struct {
	apiClient *APIClient
}

func NewSwitchMTopReportsService(c *APIClient) *SwitchMTopReportsService {
	s := new(SwitchMTopReportsService)
	s.apiClient = c
	return s
}

func (ss *SwitchMService) SwitchMTopReportsService() *SwitchMTopReportsService {
	return NewSwitchMTopReportsService(ss.apiClient)
}

// AddSwitchTopByFirmware
//
// Use this API command to retrieves top N switch count based on firmware version.
//
// Request Body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
func (s *SwitchMTopReportsService) AddSwitchTopByFirmware(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet) (*SwitchMSwitchTopSwitchesByFirmwareQueryResultList, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// AddSwitchTopByModel
//
// Use this API command to retrieve top N switch count based on switch model.
//
// Request Body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
func (s *SwitchMTopReportsService) AddSwitchTopByModel(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet) (*SwitchMSwitchTopSwitchesByModelQueryResultList, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}
