package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/splittunnel"
)

type WSGSplitTunnelProfileService struct {
	apiClient *APIClient
}

func NewWSGSplitTunnelProfileService(c *APIClient) *WSGSplitTunnelProfileService {
	s := new(WSGSplitTunnelProfileService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGSplitTunnelProfileService() *WSGSplitTunnelProfileService {
	serv := new(WSGSplitTunnelProfileService)
	serv.apiClient = ss.apiClient
	return serv
}

// AddRkszonesSplitTunnelProfilesByZoneId
//
// Create a split tunnel profile.
//
// Request Body:
//	 - body *splittunnel.CreateSplitTunnelProfile
//
// Path Parameters:
// - pZoneId string
//		- required
func (s *WSGSplitTunnelProfileService) AddRkszonesSplitTunnelProfilesByZoneId(ctx context.Context, body *splittunnel.CreateSplitTunnelProfile, pZoneId string) (*common.CreateResult, error) {
}

// DeleteRkszonesSplitTunnelProfilesById
//
// Use this API command to delete a split tunnel profile by ID.
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGSplitTunnelProfileService) DeleteRkszonesSplitTunnelProfilesById(ctx context.Context, pId string, pZoneId string) (*common.EmptyResult, error) {
}

// FindRkszonesSplitTunnelProfilesById
//
// Get a split tunnel profile by ID.
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGSplitTunnelProfileService) FindRkszonesSplitTunnelProfilesById(ctx context.Context, pId string, pZoneId string) (*splittunnel.SplitTunnelProfile, error) {
}

// FindRkszonesSplitTunnelProfilesByQueryCriteria
//
// Use this API command to retrieve a list of split tunnel profile by query criteria.
//
// Request Body:
//	 - body *common.QueryCriteriaSuperSet
func (s *WSGSplitTunnelProfileService) FindRkszonesSplitTunnelProfilesByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) (*splittunnel.SplitTunnelProfileQuery, error) {
}

// FindRkszonesSplitTunnelProfilesByZoneId
//
// Get a ID list of split tunnel profile in this Zone.
//
// Path Parameters:
// - pZoneId string
//		- required
func (s *WSGSplitTunnelProfileService) FindRkszonesSplitTunnelProfilesByZoneId(ctx context.Context, pZoneId string) (*splittunnel.SplitTunnelProfileList, error) {
}

// PartialUpdateRkszonesSplitTunnelProfilesById
//
// Use this API command to modify a split tunnel profile.
//
// Request Body:
//	 - body *splittunnel.ModifySplitTunnelProfile
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGSplitTunnelProfileService) PartialUpdateRkszonesSplitTunnelProfilesById(ctx context.Context, body *splittunnel.ModifySplitTunnelProfile, pId string, pZoneId string) (*common.EmptyResult, error) {
}

// UpdateRkszonesSplitTunnelProfilesById
//
// Use this API command to modify entire information of a split tunnel profile.
//
// Request Body:
//	 - body *splittunnel.CreateSplitTunnelProfile
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGSplitTunnelProfileService) UpdateRkszonesSplitTunnelProfilesById(ctx context.Context, body *splittunnel.CreateSplitTunnelProfile, pId string, pZoneId string) (*common.EmptyResult, error) {
}
