package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/switchm/configbackup"
)

type SwitchMSwitchConfigurationService struct {
	client *Client
}

func NewSwitchMSwitchConfigurationService(client *Client) *SwitchMSwitchConfigurationService {
	s := new(SwitchMSwitchConfigurationService)
	s.client = client
	return s
}

func (ss *SwitchMService) SwitchMSwitchConfigurationService() *SwitchMSwitchConfigurationService {
	serv := new(SwitchMSwitchConfigurationService)
	serv.client = ss.client
	return serv
}

// AddSwitchconfig
//
// Use this API command to retrieve configuration backup list with specified filters.
//
// Request Body:
//	 - body *configbackup.QueryCriteria
func (s *SwitchMSwitchConfigurationService) AddSwitchconfig(ctx context.Context, body *configbackup.QueryCriteria) (*configbackup.List, error) {
}

// AddSwitchconfigBackupDiff
//
// Use this API command to diff between two config back up files for a switch.
//
// Request Body:
//	 - body *configbackup.ConfigBackupDiffInput
func (s *SwitchMSwitchConfigurationService) AddSwitchconfigBackupDiff(ctx context.Context, body *configbackup.ConfigBackupDiffInput) (*configbackup.ConfigBackupDiff, error) {
}

// FindSwitchconfigByConfigId
//
// Use this API command to retrieve configuration backup content.
//
// Path Parameters:
// - pConfigId string
//		- required
func (s *SwitchMSwitchConfigurationService) FindSwitchconfigByConfigId(ctx context.Context, pConfigId string) error {
}

// FindSwitchconfigDownloadByConfigId
//
// Use this API command to download configuration backup content as plain text.
//
// Path Parameters:
// - pConfigId string
//		- required
func (s *SwitchMSwitchConfigurationService) FindSwitchconfigDownloadByConfigId(ctx context.Context, pConfigId string) error {
}

// UpdateSwitchconfigBackup
//
// Use this API command to backup configuration for a list of switches.
//
// Request Body:
//	 - body configbackup.SwitchIds
func (s *SwitchMSwitchConfigurationService) UpdateSwitchconfigBackup(ctx context.Context, body configbackup.SwitchIds) error {
}

// UpdateSwitchconfigBackupByGroupId
//
// Use this API command to backup configurations for all switches under a group.
//
// Path Parameters:
// - pGroupId string
//		- required
// - pGroupType string
//		- required
func (s *SwitchMSwitchConfigurationService) UpdateSwitchconfigBackupByGroupId(ctx context.Context, pGroupId string, pGroupType string) error {
}

// UpdateSwitchconfigBackupRestoreByBackupId
//
// Restore a configuration backup to the switch.
//
// Path Parameters:
// - pBackupId string
//		- required
func (s *SwitchMSwitchConfigurationService) UpdateSwitchconfigBackupRestoreByBackupId(ctx context.Context, pBackupId string) error {
}
