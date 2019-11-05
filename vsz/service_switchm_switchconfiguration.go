package vsz

// API Version: v8_1

import (
	"context"
	"errors"
	"fmt"
	"github.com/myENA/ruckus-client/vsz/types/switchm/configbackup"
	"gopkg.in/go-playground/validator.v9"
)

type SwitchMSwitchConfigurationService struct {
	apiClient *APIClient
	validate  *validator.Validate
}

func NewSwitchMSwitchConfigurationService(c *APIClient) *SwitchMSwitchConfigurationService {
	s := new(SwitchMSwitchConfigurationService)
	s.apiClient = c
	s.validate = validator.New()
	return s
}

func (ss *SwitchMService) SwitchMSwitchConfigurationService() *SwitchMSwitchConfigurationService {
	return NewSwitchMSwitchConfigurationService(ss.apiClient)
}

// AddSwitchconfig
//
// Use this API command to retrieve configuration backup list with specified filters.
//
// Request Body:
//	 - body *configbackup.QueryCriteria
func (s *SwitchMSwitchConfigurationService) AddSwitchconfig(ctx context.Context, body *configbackup.QueryCriteria) (*configbackup.List, error) {
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

// AddSwitchconfigBackup
//
// Use this API command to backup configuration for a list of switches.
//
// Request Body:
//	 - body configbackup.SwitchIds
func (s *SwitchMSwitchConfigurationService) AddSwitchconfigBackup(ctx context.Context, body configbackup.SwitchIds) (*configbackup.CreateBackupResultList, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
	if len(body) == 0 {
		return nil, errors.New("body cannot be empty")
	}
}

// AddSwitchconfigBackupDiff
//
// Use this API command to diff between two config back up files for a switch.
//
// Request Body:
//	 - body *configbackup.ConfigBackupDiffInput
func (s *SwitchMSwitchConfigurationService) AddSwitchconfigBackupDiff(ctx context.Context, body *configbackup.ConfigBackupDiffInput) (*configbackup.ConfigBackupDiff, error) {
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

// DeleteSwitchconfig
//
// Use this API command to delete config backups by a list of config backup id.
//
// Request Body:
//	 - body configbackup.BackupIds
func (s *SwitchMSwitchConfigurationService) DeleteSwitchconfig(ctx context.Context, body configbackup.BackupIds) error {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
	if len(body) == 0 {
		return errors.New("body cannot be empty")
	}
}

// DeleteSwitchconfigByConfigId
//
// Use this API command to delete the configuration backup.
//
// Path Parameters:
// - pConfigId string
//		- required
func (s *SwitchMSwitchConfigurationService) DeleteSwitchconfigByConfigId(ctx context.Context, pConfigId string) error {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
}

// FindSwitchconfigByConfigId
//
// Use this API command to retrieve configuration backup content.
//
// Path Parameters:
// - pConfigId string
//		- required
func (s *SwitchMSwitchConfigurationService) FindSwitchconfigByConfigId(ctx context.Context, pConfigId string) error {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
}

// FindSwitchconfigDownloadByConfigId
//
// Use this API command to download configuration backup content as plain text.
//
// Path Parameters:
// - pConfigId string
//		- required
func (s *SwitchMSwitchConfigurationService) FindSwitchconfigDownloadByConfigId(ctx context.Context, pConfigId string) error {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
}

// UpdateSwitchconfigBackup
//
// Use this API command to backup configuration for a list of switches.
//
// Request Body:
//	 - body configbackup.SwitchIds
func (s *SwitchMSwitchConfigurationService) UpdateSwitchconfigBackup(ctx context.Context, body configbackup.SwitchIds) error {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
	if len(body) == 0 {
		return errors.New("body cannot be empty")
	}
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
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
}

// UpdateSwitchconfigBackupRestoreByBackupId
//
// Restore a configuration backup to the switch.
//
// Path Parameters:
// - pBackupId string
//		- required
func (s *SwitchMSwitchConfigurationService) UpdateSwitchconfigBackupRestoreByBackupId(ctx context.Context, pBackupId string) error {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
}
