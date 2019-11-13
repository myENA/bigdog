package vsz

// API Version: v8_1

import (
	"context"
	"errors"
	"fmt"
	"gopkg.in/go-playground/validator.v9"
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
// Path Parameters:
// - pId string
//		- required
func (s *WSGVDPProfileService) DeleteProfilesVdpById(ctx context.Context, pId string) (*WSGCommonEmptyResult, error) {
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
//		- nullable
// - qListSize string
//		- nullable
func (s *WSGVDPProfileService) FindProfilesVdp(ctx context.Context, qIndex string, qListSize string) (*WSGProfileList, error) {
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
func (s *WSGVDPProfileService) FindProfilesVdpById(ctx context.Context, pId string) (*WSGProfileVdpProfile, error) {
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
