package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/profile"
)

type WSGBridgeService struct {
	apiClient *APIClient
}

func NewWSGBridgeService(c *APIClient) *WSGBridgeService {
	s := new(WSGBridgeService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGBridgeService() *WSGBridgeService {
	serv := new(WSGBridgeService)
	serv.apiClient = ss.apiClient
	return serv
}

// AddProfilesBridge
//
// Use this API command to create Bridge profile.
//
// Request Body:
//	 - body *profile.CreateBridgeProfile
func (s *WSGBridgeService) AddProfilesBridge(ctx context.Context, body *profile.CreateBridgeProfile) (*common.CreateResult, error) {
}

// DeleteProfilesBridge
//
// Use this API command to delete multiple bridge profile.
//
// Request Body:
//	 - body *common.BulkDeleteRequest
func (s *WSGBridgeService) DeleteProfilesBridge(ctx context.Context, body *common.BulkDeleteRequest) (*common.EmptyResult, error) {
}

// DeleteProfilesBridgeById
//
// Use this API command to delete Bridge profile.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGBridgeService) DeleteProfilesBridgeById(ctx context.Context, pId string) error {
}

// FindProfilesBridge
//
// Use this API command to retrieve a list of Bridge profile.
func (s *WSGBridgeService) FindProfilesBridge(ctx context.Context) (*profile.ProfileList, error) {
}

// FindProfilesBridgeById
//
// Use this API command to retrieve Bridge profile by ID.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGBridgeService) FindProfilesBridgeById(ctx context.Context, pId string) (*profile.BridgeProfile, error) {
}

// FindProfilesBridgeByQueryCriteria
//
// Use this API command to query a list of Bridge profile.
//
// Request Body:
//	 - body *common.QueryCriteriaSuperSet
func (s *WSGBridgeService) FindProfilesBridgeByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) (*profile.BridgeProfileList, error) {
}

// PartialUpdateProfilesBridgeById
//
// Use this API command to modify the basic information of Bridge profile.
//
// Request Body:
//	 - body *profile.ModifyBridgeProfile
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGBridgeService) PartialUpdateProfilesBridgeById(ctx context.Context, body *profile.ModifyBridgeProfile, pId string) (*common.EmptyResult, error) {
}
