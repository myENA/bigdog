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

func (s *SwitchMSwitchFirmwareService) AddFirmwareUpload(ctx context.Context) error {
}

func (s *SwitchMSwitchFirmwareService) DeleteFirmwareByVersion(ctx context.Context, pVersion string) error {
}

func (s *SwitchMSwitchFirmwareService) FindFirmware(ctx context.Context) (*firmware.AllFirmwaresQueryResultList, error) {
}
