package vsz

// API Version: v8_1

import (
	"context"
	"errors"
	"fmt"
	"gopkg.in/go-playground/validator.v9"
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
	return NewWSGDataPlaneOperationalService(ss.apiClient)
}

// AddDpsSwitchoverCluster
//
// Use this API command to switchover DP to another cluster
//
// Request Body:
//	 - body *WSGDPSwitchoverDp
func (s *WSGDataPlaneOperationalService) AddDpsSwitchoverCluster(ctx context.Context, body *WSGDPSwitchoverDp) (*WSGDPEmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}
