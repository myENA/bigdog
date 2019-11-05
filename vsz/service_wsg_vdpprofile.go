package vsz

// API Version: v8_1

import (
	"context"
	"errors"
	"fmt"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/profile"
	"gopkg.in/go-playground/validator.v9"
)

type WSGVDPProfileService struct {
	apiClient *APIClient
	validate  *validator.Validate
}

func NewWSGVDPProfileService(c *APIClient) *WSGVDPProfileService {
	s := new(WSGVDPProfileService)
	s.apiClient = c
	s.validate = validator.New()
	return s
}

func (ss *WSGService) WSGVDPProfileService() *WSGVDPProfileService {
	return NewWSGVDPProfileService(ss.apiClient)
}

// DeleteProfilesVdpById
//
// Use this API command to delete an vdp.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGVDPProfileService) DeleteProfilesVdpById(ctx context.Context, pId string) (*common.EmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindProfilesVdp
//
// Use this API command to retrieve a list of vdp.
//
// Query Parameters:
// - qIndex string
// - qListSize string
func (s *WSGVDPProfileService) FindProfilesVdp(ctx context.Context, qIndex string, qListSize string) (*profile.ProfileList, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindProfilesVdpById
//
// Use this API command to retrieve an vdp.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGVDPProfileService) FindProfilesVdpById(ctx context.Context, pId string) (*profile.VdpProfile, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// UpdateProfilesVdpApproveById
//
// Use this API command to approve vdp.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGVDPProfileService) UpdateProfilesVdpApproveById(ctx context.Context, pId string) error {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
}
