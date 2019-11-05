package vsz

// API Version: v8_1

import (
	"context"
	"errors"
	"fmt"
	"github.com/myENA/ruckus-client/vsz/types/wsg/dpprofile"
	"gopkg.in/go-playground/validator.v9"
)

type WSGDPDHCPNATProfileService struct {
	apiClient *APIClient
	validate  *validator.Validate
}

func NewWSGDPDHCPNATProfileService(c *APIClient) *WSGDPDHCPNATProfileService {
	s := new(WSGDPDHCPNATProfileService)
	s.apiClient = c
	s.validate = validator.New()
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
//	 - body *dpprofile.DpProfileSettingBO
func (s *WSGDPDHCPNATProfileService) AddDpProfileSettings(ctx context.Context, body *dpprofile.DpProfileSettingBO) error {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
	if body == nil {
		return errors.New("body cannot be empty")
	}
}

// DeleteDpProfileSettings
//
// Use this API command to delete DP DHCP & NAT profile settings.
//
// Request Body:
//	 - body *dpprofile.BulkDelete
func (s *WSGDPDHCPNATProfileService) DeleteDpProfileSettings(ctx context.Context, body *dpprofile.BulkDelete) error {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
	if body == nil {
		return errors.New("body cannot be empty")
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
func (s *WSGDPDHCPNATProfileService) FindDpProfileSettings(ctx context.Context) (*dpprofile.DpProfileSettingBOList, error) {
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
func (s *WSGDPDHCPNATProfileService) FindDpProfileSettingsByDpKey(ctx context.Context, pDpKey string) (*dpprofile.DpProfileSettingBO, error) {
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
//	 - body *dpprofile.DpProfileSettingBO
//
// Path Parameters:
// - pDpKey string
//		- required
func (s *WSGDPDHCPNATProfileService) UpdateDpProfileSettingsByDpKey(ctx context.Context, body *dpprofile.DpProfileSettingBO, pDpKey string) error {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
	if body == nil {
		return errors.New("body cannot be empty")
	}
}
