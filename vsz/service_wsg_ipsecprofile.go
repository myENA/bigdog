package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/profile"
)

type WSGIPSECProfileService struct {
	client *Client
}

func NewWSGIPSECProfileService(client *Client) *WSGIPSECProfileService {
	s := new(WSGIPSECProfileService)
	s.client = client
	return s
}

func (ss *WSGService) WSGIPSECProfileService() *WSGIPSECProfileService {
	serv := new(WSGIPSECProfileService)
	serv.client = ss.client
	return serv
}

// FindProfilesTunnelIpsec
//
// Retrieve a list of IPSEC.
//
// Query Parameters:
// - qIndex string
// - qListSize string
func (s *WSGIPSECProfileService) FindProfilesTunnelIpsec(ctx context.Context, qIndex string, qListSize string) (*profile.ProfileList, error) {
}

// FindProfilesTunnelIpsecById
//
// Retrieve a IPSEC.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGIPSECProfileService) FindProfilesTunnelIpsecById(ctx context.Context, pId string) (*profile.IpsecProfile, error) {
}

// FindProfilesTunnelIpsecByQueryCriteria
//
// Query a list of IPSEC.
//
// Request Body:
//	 - body *common.QueryCriteriaSuperSet
func (s *WSGIPSECProfileService) FindProfilesTunnelIpsecByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) (*profile.IpsecProfileList, error) {
}
