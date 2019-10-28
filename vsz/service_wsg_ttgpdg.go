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

// DeleteProfilesTtgpdgApnRealmsById
//
// Use this API command to disable the APN realm of TTG PDG profile.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGTTGPDGService) DeleteProfilesTtgpdgApnRealmsById(ctx context.Context, pId string) (*common.EmptyResult, error) {
}

// DeleteProfilesTtgpdgDhcpRelayById
//
// Use this API command to disable the DHCP relay of TTG PDG profile.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGTTGPDGService) DeleteProfilesTtgpdgDhcpRelayById(ctx context.Context, pId string) (*common.EmptyResult, error) {
}

// FindProfilesTtgpdg
//
// Use this API command to retrieve a list of TTG+PDG profile.
func (s *WSGTTGPDGService) FindProfilesTtgpdg(ctx context.Context) (*profile.ProfileList, error) {
}

// FindProfilesTtgpdgById
//
// Use this API command to retrieve TTG+PDG profile by ID.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGTTGPDGService) FindProfilesTtgpdgById(ctx context.Context, pId string) (*profile.TtgpdgProfile, error) {
}

// FindProfilesTtgpdgByQueryCriteria
//
// Use this API command to query a list of TTG+PDG profile.
//
// Request Body:
//	 - body *common.QueryCriteriaSuperSet
func (s *WSGTTGPDGService) FindProfilesTtgpdgByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) (*profile.TtgpdgProfileList, error) {
}
