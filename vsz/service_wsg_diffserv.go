package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/zone"
)

type WSGDiffServService struct {
    client *Client
}

func NewWSGDiffServService (client *Client) *WSGDiffServService {
    s := new(WSGDiffServService)
    s.client = client
    return s
}

func (ss *WSGService) WSGDiffServService () *WSGDiffServService {
    serv := new(WSGDiffServService)
    serv.client = ss.client
    return serv
}

func (s *WSGDiffServService) FindRkszonesDiffservById (ctx context.Context, id string, zoneId string) (*zone.DiffServConfiguration, error) {
}

func (s *WSGDiffServService) FindRkszonesDiffservByZoneId (ctx context.Context, zoneId string) (*zone.DiffServList, error) {
}

