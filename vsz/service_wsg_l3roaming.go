package vsz

// API Version: v8_0

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/profile"
)

type WSGL3RoamingService struct {
    client *Client
}

func NewWSGL3RoamingService (client *Client) *WSGL3RoamingService {
    s := new(WSGL3RoamingService)
    s.client = client
    return s
}

func (ss *WSGService) WSGL3RoamingService () *WSGL3RoamingService {
    serv := new(WSGL3RoamingService)
    serv.client = ss.client
    return serv
}

func (s *WSGL3RoamingService) FindProfilesTunnelL3Roaming (ctx context.Context) (profile.GetL3RoamingConfig, error) {
}

func (s *WSGL3RoamingService) FindProfilesTunnelL3RoamingByQueryCriteria (ctx context.Context) (profile.GetL3RoamingConfig, error) {
}

func (s *WSGL3RoamingService) PartialUpdateProfilesTunnelL3RoamingDataPlanes (ctx context.Context) error {
}

