package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/dp"
)

type WSGDataPlaneOperationalService struct {
	apiClient *APIClient
}

func NewWSGDataPlaneOperationalService(c *APIClient) *WSGDataPlaneOperationalService {
	s := new(WSGDataPlaneOperationalService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGDataPlaneOperationalService() *WSGDataPlaneOperationalService {
	serv := new(WSGDataPlaneOperationalService)
	serv.apiClient = ss.apiClient
	return serv
}

// AddDpsSwitchoverCluster
//
// Use this API command to switchover DP to another cluster
//
// Request Body:
//	 - body *dp.SwitchoverDp
func (s *WSGDataPlaneOperationalService) AddDpsSwitchoverCluster(ctx context.Context, body *dp.SwitchoverDp) (*dp.EmptyResult, error) {
}
