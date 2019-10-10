package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/profile"
)

type WSGTTGPDGService struct {
	client *Client
}

func NewWSGTTGPDGService(client *Client) *WSGTTGPDGService {
	s := new(WSGTTGPDGService)
	s.client = client
	return s
}

func (ss *WSGService) WSGTTGPDGService() *WSGTTGPDGService {
	serv := new(WSGTTGPDGService)
	serv.client = ss.client
	return serv
}

func (s *WSGTTGPDGService) DeleteProfilesTtgpdgApnRealmsById(ctx context.Context, id string) (*common.EmptyResult, error) {
}

func (s *WSGTTGPDGService) DeleteProfilesTtgpdgDhcpRelayById(ctx context.Context, id string) (*common.EmptyResult, error) {
}

func (s *WSGTTGPDGService) FindProfilesTtgpdg(ctx context.Context) (*profile.ProfileList, error) {
}

func (s *WSGTTGPDGService) FindProfilesTtgpdgById(ctx context.Context, id string) (*profile.TtgpdgProfile, error) {
}

func (s *WSGTTGPDGService) FindProfilesTtgpdgByQueryCriteria(ctx context.Context) (*profile.TtgpdgProfileList, error) {
}
