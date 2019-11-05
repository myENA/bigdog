package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/switchm/common"
	"github.com/myENA/ruckus-client/vsz/types/switchm/firmware"
)

type SwitchMSwitchFirmwareService struct {
	apiClient *APIClient
}

func NewSwitchMSwitchFirmwareService(c *APIClient) *SwitchMSwitchFirmwareService {
	s := new(SwitchMSwitchFirmwareService)
	s.apiClient = c
	return s
}

func (ss *SwitchMService) SwitchMSwitchFirmwareService() *SwitchMSwitchFirmwareService {
	serv := new(SwitchMSwitchFirmwareService)
	serv.apiClient = ss.apiClient
	return serv
}

// AddFirmware
//
// Use this API command to retrieve list of switch firmwares uploaded to SmartZone.
//
// Request Body:
//	 - body *common.QueryCriteriaSuperSet
func (s *SwitchMSwitchFirmwareService) AddFirmware(ctx context.Context, body *common.QueryCriteriaSuperSet) (*firmware.FirmwaresQueryResultList, error) {
}

// AddFirmwareUpload
//
// Use this API command to upload a firmware image zip file to SmartZone.
func (s *SwitchMSwitchFirmwareService) AddFirmwareUpload(ctx context.Context) error {
}

// DeleteFirmwareByVersion
//
// Use this API command to deletes a firmware image file from SmartZone.
//
// Path Parameters:
// - pVersion string
//		- required
func (s *SwitchMSwitchFirmwareService) DeleteFirmwareByVersion(ctx context.Context, pVersion string) error {
}

// FindFirmware
//
// Use this API command to retrieve list of switch firmwares uploaded to SmartZone.
func (s *SwitchMSwitchFirmwareService) FindFirmware(ctx context.Context) (*firmware.AllFirmwaresQueryResultList, error) {
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
}
