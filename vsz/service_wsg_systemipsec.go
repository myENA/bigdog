package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/systemipsec"
)

type WSGSystemIPsecService struct {
	client *Client
}

func NewWSGSystemIPsecService(client *Client) *WSGSystemIPsecService {
	s := new(WSGSystemIPsecService)
	s.client = client
	return s
}

func (ss *WSGService) WSGSystemIPsecService() *WSGSystemIPsecService {
	serv := new(WSGSystemIPsecService)
	serv.client = ss.client
	return serv
}

func (s *WSGSystemIPsecService) FindSystemIpsec(ctx context.Context) (*systemipsec.GetResult, error) {
}
