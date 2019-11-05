package vsz

// API Version: v8_1

import (
	"context"
	"errors"
	"fmt"
	"github.com/myENA/ruckus-client/vsz/types/switchm/common"
	"github.com/myENA/ruckus-client/vsz/types/switchm/firmware"
	"gopkg.in/go-playground/validator.v9"
)

type SwitchMSwitchFirmwareService struct {
	apiClient *APIClient
	validate  *validator.Validate
}

func NewSwitchMSwitchFirmwareService(c *APIClient) *SwitchMSwitchFirmwareService {
	s := new(SwitchMSwitchFirmwareService)
	s.apiClient = c
	s.validate = validator.New()
	return s
}

func (ss *SwitchMService) SwitchMSwitchFirmwareService() *SwitchMSwitchFirmwareService {
	return NewSwitchMSwitchFirmwareService(ss.apiClient)
}

// AddFirmware
//
// Use this API command to retrieve list of switch firmwares uploaded to SmartZone.
//
// Request Body:
//	 - body *common.QueryCriteriaSuperSet
func (s *SwitchMSwitchFirmwareService) AddFirmware(ctx context.Context, body *common.QueryCriteriaSuperSet) (*firmware.FirmwaresQueryResultList, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
	if body == nil {
		return nil, errors.New("body cannot be empty")
	}
}

// AddFirmwareUpload
//
// Use this API command to upload a firmware image zip file to SmartZone.
func (s *SwitchMSwitchFirmwareService) AddFirmwareUpload(ctx context.Context) error {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteFirmwareByVersion
//
// Use this API command to deletes a firmware image file from SmartZone.
//
// Path Parameters:
// - pVersion string
//		- required
func (s *SwitchMSwitchFirmwareService) DeleteFirmwareByVersion(ctx context.Context, pVersion string) error {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
}

// FindFirmware
//
// Use this API command to retrieve list of switch firmwares uploaded to SmartZone.
func (s *SwitchMSwitchFirmwareService) FindFirmware(ctx context.Context) (*firmware.AllFirmwaresQueryResultList, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// PartialUpdateFirmwareByVersion
//
// Use this API command to update the given firmware version on switches matching criteria.
//
// Request Body:
//	 - body *common.QueryCriteriaSuperSet
//
// Path Parameters:
// - pVersion string
//		- required
func (s *SwitchMSwitchFirmwareService) PartialUpdateFirmwareByVersion(ctx context.Context, body *common.QueryCriteriaSuperSet, pVersion string) (*firmware.ScheduleIds, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
	if body == nil {
		return nil, errors.New("body cannot be empty")
	}
}
