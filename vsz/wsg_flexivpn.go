package vsz

// API Version: v8_1

import (
	"context"
	"errors"
	"fmt"
	"gopkg.in/go-playground/validator.v9"
)

type WSGFlexiVPNService struct {
	apiClient *APIClient
}

func NewWSGFlexiVPNService(c *APIClient) *WSGFlexiVPNService {
	s := new(WSGFlexiVPNService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGFlexiVPNService() *WSGFlexiVPNService {
	return NewWSGFlexiVPNService(ss.apiClient)
}

type WSGFlexiVPNSetting struct {
	// ZoneAffinityId
	// Zone Affinity ID
	// Constraints:
	//    - required
	ZoneAffinityId *string `json:"zoneAffinityId" validate:"required"`
}

// DeleteRkszonesWlansFlexiVpnProfileById
//
// Use this API command to delete Flexi-VPN on WLAN
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGFlexiVPNService) DeleteRkszonesWlansFlexiVpnProfileById(ctx context.Context, pId string, pZoneId string) error {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
}

// FindServicesFlexiVpnProfileByQueryCriteria
//
// Use this API command to query Flexi-VPN profiles.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGFlexiVPNService) FindServicesFlexiVpnProfileByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet) (*WSGProfileFlexiVpnProfileList, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}
