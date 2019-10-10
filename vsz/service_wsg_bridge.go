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

func (s *WSGBridgeService) FindProfilesBridge(ctx context.Context) (*profile.ProfileList, error) {
}

func (s *WSGBridgeService) FindProfilesBridgeById(ctx context.Context, id string) (*profile.BridgeProfile, error) {
}

func (s *WSGBridgeService) FindProfilesBridgeByQueryCriteria(ctx context.Context) (*profile.BridgeProfileList, error) {
}
