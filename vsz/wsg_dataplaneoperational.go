package vsz

// API Version: v8_1

import (
	"context"
	"errors"
	"fmt"
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
	var (
		resp *WSGDPEmptyResult
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
}
