package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/profile"
)

type WSGIPSECProfileService struct {
	apiClient *APIClient
}

func NewWSGIPSECProfileService(c *APIClient) *WSGIPSECProfileService {
	s := new(WSGIPSECProfileService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGIPSECProfileService() *WSGIPSECProfileService {
	serv := new(WSGIPSECProfileService)
	serv.apiClient = ss.apiClient
	return serv
}

// AddProfilesTunnelIpsec
//
// Create a new ipsec.
//
// Request Body:
//	 - body *profile.CreateIpsecProfile
func (s *WSGIPSECProfileService) AddProfilesTunnelIpsec(ctx context.Context, body *profile.CreateIpsecProfile) (*common.CreateResult, error) {
}

// DeleteProfilesTunnelIpsec
//
// Delete multiple ipsec.
//
// Request Body:
//	 - body *common.BulkDeleteRequest
func (s *WSGIPSECProfileService) DeleteProfilesTunnelIpsec(ctx context.Context, body *common.BulkDeleteRequest) (*common.EmptyResult, error) {
}

// DeleteProfilesTunnelIpsecById
//
// Delete a ipsec.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGIPSECProfileService) DeleteProfilesTunnelIpsecById(ctx context.Context, pId string) (*common.EmptyResult, error) {
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

// PartialUpdateProfilesTunnelIpsecById
//
// Modify a specific ipsec basic.
//
// Request Body:
//	 - body *profile.ModifyIpsecProfile
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGIPSECProfileService) PartialUpdateProfilesTunnelIpsecById(ctx context.Context, body *profile.ModifyIpsecProfile, pId string) (*common.EmptyResult, error) {
}
