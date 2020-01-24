package vsz

// API Version: v8_1

import (
	"context"
	"net/http"
)

type WSGDPDHCPNATProfileService struct {
	apiClient *APIClient
}

func NewWSGDPDHCPNATProfileService(c *APIClient) *WSGDPDHCPNATProfileService {
	s := new(WSGDPDHCPNATProfileService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGDPDHCPNATProfileService() *WSGDPDHCPNATProfileService {
	return NewWSGDPDHCPNATProfileService(ss.apiClient)
}

// AddDpProfileSettings
//
// Use this API command to create DP DHCP & NAT profile setting.
//
// Request Body:
//	 - body *WSGDPProfileSettingBO
func (s *WSGDPDHCPNATProfileService) AddDpProfileSettings(ctx context.Context, body *WSGDPProfileSettingBO) error {
	var err error
	if err = ctx.Err(); err != nil {
		return err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return err
	}
	req := NewAPIRequest(http.MethodPost, RouteWSGAddDpProfileSettings, true)
}

// DeleteDpProfileSettings
//
// Use this API command to delete DP DHCP & NAT profile settings.
//
// Request Body:
//	 - body *WSGDPProfileBulkDelete
func (s *WSGDPDHCPNATProfileService) DeleteDpProfileSettings(ctx context.Context, body *WSGDPProfileBulkDelete) error {
	var err error
	if err = ctx.Err(); err != nil {
		return err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return err
	}
	req := NewAPIRequest(http.MethodDelete, RouteWSGDeleteDpProfileSettings, true)
}

// DeleteDpProfileSettingsByDpKey
//
// Use this API command to delete DP DHCP & NAT profile setting.
//
// Required Parameters:
// - dpKey string
//		- required
func (s *WSGDPDHCPNATProfileService) DeleteDpProfileSettingsByDpKey(ctx context.Context, dpKey string) error {
	var err error
	if err = ctx.Err(); err != nil {
		return err
	}
	if err = pkgValidator.VarCtx(ctx, dpKey, "required"); err != nil {
		return err
	}
	req := NewAPIRequest(http.MethodDelete, RouteWSGDeleteDpProfileSettingsByDpKey, true)
}

// FindDpProfileSettings
//
// Use this API command to retrieve DP DHCP & NAT profile setting list.
func (s *WSGDPDHCPNATProfileService) FindDpProfileSettings(ctx context.Context) (*WSGDPProfileSettingBOList, error) {
	var (
		resp *WSGDPProfileSettingBOList
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	req := NewAPIRequest(http.MethodGet, RouteWSGFindDpProfileSettings, true)
}

// FindDpProfileSettingsByDpKey
//
// Use this API command to retrieve DP DHCP & NAT profile setting.
//
// Required Parameters:
// - dpKey string
//		- required
func (s *WSGDPDHCPNATProfileService) FindDpProfileSettingsByDpKey(ctx context.Context, dpKey string) (*WSGDPProfileSettingBO, error) {
	var (
		resp *WSGDPProfileSettingBO
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, dpKey, "required"); err != nil {
		return resp, err
	}
	req := NewAPIRequest(http.MethodGet, RouteWSGFindDpProfileSettingsByDpKey, true)
}

// UpdateDpProfileSettingsByDpKey
//
// Use this API command to modify DP DHCP & NAT profile setting.
//
// Request Body:
//	 - body *WSGDPProfileSettingBO
//
// Required Parameters:
// - dpKey string
//		- required
func (s *WSGDPDHCPNATProfileService) UpdateDpProfileSettingsByDpKey(ctx context.Context, body *WSGDPProfileSettingBO, dpKey string) error {
	var err error
	if err = ctx.Err(); err != nil {
		return err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return err
	}
	if err = pkgValidator.VarCtx(ctx, dpKey, "required"); err != nil {
		return err
	}
	req := NewAPIRequest(http.MethodPut, RouteWSGUpdateDpProfileSettingsByDpKey, true)
}
