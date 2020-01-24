package vsz

// API Version: v8_1

import (
	"context"
	"net/http"
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
	var (
		resp *SwitchMConfigBackupList
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
	req := NewAPIRequest(http.MethodPost, RouteSwitchMAddSwitchconfig, true)
}

// AddSwitchconfigBackup
//
// Use this API command to backup configuration for a list of switches.
//
// Request Body:
//	 - body SwitchMConfigBackupSwitchIds
func (s *SwitchMConfigurationService) AddSwitchconfigBackup(ctx context.Context, body SwitchMConfigBackupSwitchIds) (*SwitchMConfigBackupCreateBackupResultList, error) {
	var (
		resp *SwitchMConfigBackupCreateBackupResultList
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	}
	req := NewAPIRequest(http.MethodPost, RouteSwitchMAddSwitchconfigBackup, true)
}

// AddSwitchconfigBackupDiff
//
// Use this API command to diff between two config back up files for a switch.
//
// Request Body:
//	 - body *SwitchMConfigBackupDiffInput
func (s *SwitchMConfigurationService) AddSwitchconfigBackupDiff(ctx context.Context, body *SwitchMConfigBackupDiffInput) (*SwitchMConfigBackupDiff, error) {
	var (
		resp *SwitchMConfigBackupDiff
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
	req := NewAPIRequest(http.MethodPost, RouteSwitchMAddSwitchconfigBackupDiff, true)
}

// DeleteSwitchconfig
//
// Use this API command to delete config backups by a list of config backup id.
//
// Request Body:
//	 - body SwitchMConfigBackupBackupIds
func (s *SwitchMConfigurationService) DeleteSwitchconfig(ctx context.Context, body SwitchMConfigBackupBackupIds) error {
	var err error
	if err = ctx.Err(); err != nil {
		return err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return err
	}
	req := NewAPIRequest(http.MethodDelete, RouteSwitchMDeleteSwitchconfig, true)
}

// DeleteSwitchconfigByConfigId
//
// Use this API command to delete the configuration backup.
//
// Required Parameters:
// - configId string
//		- required
func (s *SwitchMConfigurationService) DeleteSwitchconfigByConfigId(ctx context.Context, configId string) error {
	var err error
	if err = ctx.Err(); err != nil {
		return err
	}
	if err = pkgValidator.VarCtx(ctx, configId, "required"); err != nil {
		return err
	}
	req := NewAPIRequest(http.MethodDelete, RouteSwitchMDeleteSwitchconfigByConfigId, true)
}

// FindSwitchconfigByConfigId
//
// Use this API command to retrieve configuration backup content.
//
// Required Parameters:
// - configId string
//		- required
func (s *SwitchMConfigurationService) FindSwitchconfigByConfigId(ctx context.Context, configId string) (interface{}, error) {
	var (
		resp interface{}
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, configId, "required"); err != nil {
		return resp, err
	}
	req := NewAPIRequest(http.MethodGet, RouteSwitchMFindSwitchconfigByConfigId, true)
}

// FindSwitchconfigDownloadByConfigId
//
// Use this API command to download configuration backup content as plain text.
//
// Required Parameters:
// - configId string
//		- required
func (s *SwitchMConfigurationService) FindSwitchconfigDownloadByConfigId(ctx context.Context, configId string) (interface{}, error) {
	var (
		resp interface{}
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, configId, "required"); err != nil {
		return resp, err
	}
	req := NewAPIRequest(http.MethodGet, RouteSwitchMFindSwitchconfigDownloadByConfigId, true)
}

// UpdateSwitchconfigBackup
//
// Use this API command to backup configuration for a list of switches.
//
// Request Body:
//	 - body SwitchMConfigBackupSwitchIds
func (s *SwitchMConfigurationService) UpdateSwitchconfigBackup(ctx context.Context, body SwitchMConfigBackupSwitchIds) error {
	var err error
	if err = ctx.Err(); err != nil {
		return err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return err
	}
	req := NewAPIRequest(http.MethodPut, RouteSwitchMUpdateSwitchconfigBackup, true)
}

// UpdateSwitchconfigBackupByGroupId
//
// Use this API command to backup configurations for all switches under a group.
//
// Required Parameters:
// - groupId string
//		- required
// - groupType string
//		- required
func (s *SwitchMConfigurationService) UpdateSwitchconfigBackupByGroupId(ctx context.Context, groupId string, groupType string) error {
	var err error
	if err = ctx.Err(); err != nil {
		return err
	}
	if err = pkgValidator.VarCtx(ctx, groupType, "required"); err != nil {
		return err
	}
	if err = pkgValidator.VarCtx(ctx, groupId, "required"); err != nil {
		return err
	}
	req := NewAPIRequest(http.MethodPut, RouteSwitchMUpdateSwitchconfigBackupByGroupId, true)
}

// UpdateSwitchconfigBackupRestoreByBackupId
//
// Restore a configuration backup to the switch.
//
// Required Parameters:
// - backupId string
//		- required
func (s *SwitchMConfigurationService) UpdateSwitchconfigBackupRestoreByBackupId(ctx context.Context, backupId string) (interface{}, error) {
	var (
		resp interface{}
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, backupId, "required"); err != nil {
		return resp, err
	}
	req := NewAPIRequest(http.MethodPut, RouteSwitchMUpdateSwitchconfigBackupRestoreByBackupId, true)
}
