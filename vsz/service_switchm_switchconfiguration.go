package vsz

// API Version: v8_0

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/switchm/configbackup"
)

type SwitchMSwitchConfigurationService struct {
    client *Client
}

func NewSwitchMSwitchConfigurationService (client *Client) *SwitchMSwitchConfigurationService {
    s := new(SwitchMSwitchConfigurationService)
    s.client = client
    return s
}

func (ss *SwitchMService) SwitchMSwitchConfigurationService () *SwitchMSwitchConfigurationService {
    serv := new(SwitchMSwitchConfigurationService)
    serv.client = ss.client
    return serv
}

func (s *SwitchMSwitchConfigurationService) AddSwitchconfig (ctx context.Context) (configbackup.List, error) {
}

func (s *SwitchMSwitchConfigurationService) AddSwitchconfigBackupDiff (ctx context.Context) (configbackup.ConfigBackupDiff, error) {
}

func (s *SwitchMSwitchConfigurationService) FindSwitchconfigByConfigId (ctx context.Context, configId string) error {
}

func (s *SwitchMSwitchConfigurationService) FindSwitchconfigDownloadByConfigId (ctx context.Context, configId string) error {
}

func (s *SwitchMSwitchConfigurationService) UpdateSwitchconfigBackup (ctx context.Context) error {
}

func (s *SwitchMSwitchConfigurationService) UpdateSwitchconfigBackupByGroupId (ctx context.Context, groupId string, groupType string) error {
}

func (s *SwitchMSwitchConfigurationService) UpdateSwitchconfigBackupRestoreByBackupId (ctx context.Context, backupId string) error {
}

