package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/profile"
)

type WSGBridgeService struct {
	client *Client
}

func NewWSGBridgeService(client *Client) *WSGBridgeService {
	s := new(WSGBridgeService)
	s.client = client
	return s
}

func (ss *WSGService) WSGBridgeService() *WSGBridgeService {
	serv := new(WSGBridgeService)
	serv.client = ss.client
	return serv
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
