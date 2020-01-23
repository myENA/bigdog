package vsz

// API Version: v8_1

import (
	"context"
	"errors"
	"fmt"
)

type WSGVDPProfileService struct {
	apiClient *APIClient
}

func NewWSGVDPProfileService(c *APIClient) *WSGVDPProfileService {
	s := new(WSGVDPProfileService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGVDPProfileService() *WSGVDPProfileService {
	return NewWSGVDPProfileService(ss.apiClient)
}

// DeleteProfilesVdpById
//
// Use this API command to delete an vdp.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGVDPProfileService) DeleteProfilesVdpById(ctx context.Context, id string) (*WSGCommonEmptyResult, error) {
	var (
		resp *WSGCommonEmptyResult
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
}

// FindProfilesVdp
//
// Use this API command to retrieve a list of vdp.
//
// Optional Parameters:
// - index string
//		- nullable
// - listSize string
//		- nullable
func (s *WSGVDPProfileService) FindProfilesVdp(ctx context.Context, optionalParams map[string]interface{}) (*WSGProfileList, error) {
	var (
		resp *WSGProfileList
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
}

// FindProfilesVdpById
//
// Use this API command to retrieve an vdp.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGVDPProfileService) FindProfilesVdpById(ctx context.Context, id string) (*WSGProfileVdpProfile, error) {
	var (
		resp *WSGProfileVdpProfile
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
}

// UpdateProfilesVdpApproveById
//
// Use this API command to approve vdp.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGVDPProfileService) UpdateProfilesVdpApproveById(ctx context.Context, id string) error {
	var err error
	if err = ctx.Err(); err != nil {
		return err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return err
	}
}
