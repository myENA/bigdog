package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/switchm/firmware"
)

type SwitchMSwitchFirmwareService struct {
	client *Client
}

func NewSwitchMSwitchFirmwareService(client *Client) *SwitchMSwitchFirmwareService {
	s := new(SwitchMSwitchFirmwareService)
	s.client = client
	return s
}

func (ss *SwitchMService) SwitchMSwitchFirmwareService() *SwitchMSwitchFirmwareService {
	serv := new(SwitchMSwitchFirmwareService)
	serv.client = ss.client
	return serv
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
