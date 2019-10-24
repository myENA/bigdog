package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/portalservice"
)

type WSGWechatService struct {
	client *Client
}

func NewWSGWechatService(client *Client) *WSGWechatService {
	s := new(WSGWechatService)
	s.client = client
	return s
}

func (ss *WSGService) WSGWechatService() *WSGWechatService {
	serv := new(WSGWechatService)
	serv.client = ss.client
	return serv
}

// FindRkszonesPortalsWechatById
//
// Use this API command to retrieve wechat portal by ID.
func (s *WSGWechatService) FindRkszonesPortalsWechatById(ctx context.Context, pId string, pZoneId string) (*portalservice.WechatConfiguration, error) {
}

// FindRkszonesPortalsWechatByZoneId
//
// Use this API command to retrieve a list of wechat portal.
func (s *WSGWechatService) FindRkszonesPortalsWechatByZoneId(ctx context.Context, pZoneId string, qIndex string, qListSize string) (*portalservice.PortalServiceList, error) {
}
