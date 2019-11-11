package vsz

// API Version: v8_1

import (
	"context"
	"errors"
	"fmt"
	"gopkg.in/go-playground/validator.v9"
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
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteDpProfileSettings
//
// Use this API command to delete DP DHCP & NAT profile settings.
//
// Request Body:
//	 - body *WSGDPProfileBulkDelete
func (s *WSGDPDHCPNATProfileService) DeleteDpProfileSettings(ctx context.Context, body *WSGDPProfileBulkDelete) error {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteDpProfileSettingsByDpKey
//
// Use this API command to delete DP DHCP & NAT profile setting.
//
// Path Parameters:
// - pDpKey string
//		- required
func (s *WSGDPDHCPNATProfileService) DeleteDpProfileSettingsByDpKey(ctx context.Context, pDpKey string) error {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
}

// FindDpProfileSettings
//
// Use this API command to retrieve DP DHCP & NAT profile setting list.
func (s *WSGDPDHCPNATProfileService) FindDpProfileSettings(ctx context.Context) (*WSGDPProfileSettingBOList, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindDpProfileSettingsByDpKey
//
// Use this API command to retrieve DP DHCP & NAT profile setting.
//
// Path Parameters:
// - pDpKey string
//		- required
func (s *WSGDPDHCPNATProfileService) FindDpProfileSettingsByDpKey(ctx context.Context, pDpKey string) (*WSGDPProfileSettingBO, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// UpdateDpProfileSettingsByDpKey
//
// Use this API command to modify DP DHCP & NAT profile setting.
//
// Request Body:
//	 - body *WSGDPProfileSettingBO
//
// Path Parameters:
// - pDpKey string
//		- required
func (s *WSGDPDHCPNATProfileService) UpdateDpProfileSettingsByDpKey(ctx context.Context, body *WSGDPProfileSettingBO, pDpKey string) error {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
}
