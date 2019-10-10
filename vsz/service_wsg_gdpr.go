package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/gdpr"
)

type WSGGDPRService struct {
	client *Client
}

func NewWSGGDPRService(client *Client) *WSGGDPRService {
	s := new(WSGGDPRService)
	s.client = client
	return s
}

func (ss *WSGService) WSGGDPRService() *WSGGDPRService {
	serv := new(WSGGDPRService)
	serv.client = ss.client
	return serv
}

func (s *WSGGDPRService) AddGdprReport(ctx context.Context) error {
}
