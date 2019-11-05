package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/profile"
)

type WSGZoneAffinityProfileService struct {
	apiClient *APIClient
}

func NewWSGZoneAffinityProfileService(c *APIClient) *WSGZoneAffinityProfileService {
	s := new(WSGZoneAffinityProfileService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGZoneAffinityProfileService() *WSGZoneAffinityProfileService {
	serv := new(WSGZoneAffinityProfileService)
	serv.apiClient = ss.apiClient
	return serv
}

// AddProfilesZoneAffinity
//
// Use this API command to create zone affinity profile.
//
// Request Body:
//	 - body *profile.CreateZoneAffinityProfile
func (s *WSGZoneAffinityProfileService) AddProfilesZoneAffinity(ctx context.Context, body *profile.CreateZoneAffinityProfile) (*common.CreateResult, error) {
}

// DeleteProfilesZoneAffinityById
//
// Use this API command to delete zone affinity profile.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGZoneAffinityProfileService) DeleteProfilesZoneAffinityById(ctx context.Context, pId string) error {
}

// FindProfilesZoneAffinity
//
// Use this API command to get all zone affinity profiles.
//
// Query Parameters:
// - qVdpId string
func (s *WSGZoneAffinityProfileService) FindProfilesZoneAffinity(ctx context.Context, qVdpId string) (*profile.ZoneAffinityProfileList, error) {
}

// FindProfilesZoneAffinityById
//
// Use this API command to get one zone affinity profile.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGZoneAffinityProfileService) FindProfilesZoneAffinityById(ctx context.Context, pId string) (*profile.ReturnZoneAffinityProfile, error) {
}

// PartialUpdateProfilesZoneAffinityById
//
// Use this API command to modify zone affinity profile.
//
// Request Body:
//	 - body *profile.ModifyZoneAffinityProfile
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGZoneAffinityProfileService) PartialUpdateProfilesZoneAffinityById(ctx context.Context, body *profile.ModifyZoneAffinityProfile, pId string) (*common.EmptyResult, error) {
}
