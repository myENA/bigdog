package vsz

// API Version: v8_1

import (
	"context"
	"errors"
	"fmt"
	"gopkg.in/go-playground/validator.v9"
)

type SwitchMConfigurationService struct {
	apiClient *APIClient
}

func NewSwitchMConfigurationService(c *APIClient) *SwitchMConfigurationService {
	s := new(SwitchMConfigurationService)
	s.apiClient = c
	return s
}

func (ss *SwitchMService) SwitchMConfigurationService() *SwitchMConfigurationService {
	return NewSwitchMConfigurationService(ss.apiClient)
}

// AddSwitchconfig
//
// Use this API command to retrieve configuration backup list with specified filters.
//
// Request Body:
//	 - body *SwitchMConfigBackupQueryCriteria
func (s *SwitchMConfigurationService) AddSwitchconfig(ctx context.Context, body *SwitchMConfigBackupQueryCriteria) (*SwitchMConfigBackupList, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// AddSwitchconfigBackup
//
// Use this API command to backup configuration for a list of switches.
//
// Request Body:
//	 - body SwitchMConfigBackupSwitchIds
func (s *SwitchMConfigurationService) AddSwitchconfigBackup(ctx context.Context, body SwitchMConfigBackupSwitchIds) (*SwitchMConfigBackupCreateBackupResultList, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// AddSwitchconfigBackupDiff
//
// Use this API command to diff between two config back up files for a switch.
//
// Request Body:
//	 - body *SwitchMConfigBackupDiffInput
func (s *SwitchMConfigurationService) AddSwitchconfigBackupDiff(ctx context.Context, body *SwitchMConfigBackupDiffInput) (*SwitchMConfigBackupDiff, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteSwitchconfig
//
// Use this API command to delete config backups by a list of config backup id.
//
// Request Body:
//	 - body SwitchMConfigBackupBackupIds
func (s *SwitchMConfigurationService) DeleteSwitchconfig(ctx context.Context, body SwitchMConfigBackupBackupIds) error {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteSwitchconfigByConfigId
//
// Use this API command to delete the configuration backup.
//
// Path Parameters:
// - pConfigId string
//		- required
func (s *SwitchMConfigurationService) DeleteSwitchconfigByConfigId(ctx context.Context, pConfigId string) error {
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
func (s *SwitchMConfigurationService) FindSwitchconfigByConfigId(ctx context.Context, pConfigId string) error {
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
func (s *SwitchMConfigurationService) FindSwitchconfigDownloadByConfigId(ctx context.Context, pConfigId string) error {
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
//	 - body SwitchMConfigBackupSwitchIds
func (s *SwitchMConfigurationService) UpdateSwitchconfigBackup(ctx context.Context, body SwitchMConfigBackupSwitchIds) error {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
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
func (s *SwitchMConfigurationService) UpdateSwitchconfigBackupByGroupId(ctx context.Context, pGroupId string, pGroupType string) error {
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
func (s *SwitchMConfigurationService) UpdateSwitchconfigBackupRestoreByBackupId(ctx context.Context, pBackupId string) error {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
}
