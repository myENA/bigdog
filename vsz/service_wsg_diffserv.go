package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/zone"
)

type WSGDiffServService struct {
	client *Client
}

func NewWSGDiffServService(client *Client) *WSGDiffServService {
	s := new(WSGDiffServService)
	s.client = client
	return s
}

func (ss *WSGService) WSGDiffServService() *WSGDiffServService {
	serv := new(WSGDiffServService)
	serv.client = ss.client
	return serv
}

// FindRkszonesDiffservById
//
// Use this API command to retrieve DiffServ profile.
func (s *WSGDiffServService) FindRkszonesDiffservById(ctx context.Context, pId string, pZoneId string) (*zone.DiffServConfiguration, error) {
}

// FindRkszonesDiffservByZoneId
//
// Use this API command to retrieve a list of DiffServ profile.
func (s *WSGDiffServService) FindRkszonesDiffservByZoneId(ctx context.Context, pZoneId string) (*zone.DiffServList, error) {
}
