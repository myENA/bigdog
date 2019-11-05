package vsz

// API Version: v8_1

import (
	"context"
	"errors"
	"fmt"
	"github.com/myENA/ruckus-client/vsz/types/wsg/dp"
	"gopkg.in/go-playground/validator.v9"
)

type WSGDataPlaneOperationalService struct {
	apiClient *APIClient
	validate  *validator.Validate
}

func NewWSGDataPlaneOperationalService(c *APIClient) *WSGDataPlaneOperationalService {
	s := new(WSGDataPlaneOperationalService)
	s.apiClient = c
	s.validate = validator.New()
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
//	 - body *dp.SwitchoverDp
func (s *WSGDataPlaneOperationalService) AddDpsSwitchoverCluster(ctx context.Context, body *dp.SwitchoverDp) (*dp.EmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
	if body == nil {
		return nil, errors.New("body cannot be empty")
	}
}
