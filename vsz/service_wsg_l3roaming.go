package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/profile"
)

type WSGL3RoamingService struct {
	apiClient *APIClient
}

func NewWSGL3RoamingService(c *APIClient) *WSGL3RoamingService {
	s := new(WSGL3RoamingService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGL3RoamingService() *WSGL3RoamingService {
	serv := new(WSGL3RoamingService)
	serv.apiClient = ss.apiClient
	return serv
}

// FindProfilesTunnelL3Roaming
//
// Use this API command to retrieve L3 Roaming basic configuration.
func (s *WSGL3RoamingService) FindProfilesTunnelL3Roaming(ctx context.Context) (*profile.GetL3RoamingConfig, error) {
}

// FindProfilesTunnelL3RoamingByQueryCriteria
//
// Use this API command to retrieve L3 Roaming configuration.
func (s *WSGL3RoamingService) FindProfilesTunnelL3RoamingByQueryCriteria(ctx context.Context) (*profile.GetL3RoamingConfig, error) {
}

// PartialUpdateProfilesTunnelL3Roaming
//
// Use this API command to modify L3 Roaming basic configuration.
//
// Request Body:
//	 - body *profile.UpdateL3RoamingConfig
func (s *WSGL3RoamingService) PartialUpdateProfilesTunnelL3Roaming(ctx context.Context, body *profile.UpdateL3RoamingConfig) (*common.EmptyResult, error) {
}
