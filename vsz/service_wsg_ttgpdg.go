package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/profile"
)

type WSGTTGPDGService struct {
	apiClient *APIClient
}

func NewWSGTTGPDGService(c *APIClient) *WSGTTGPDGService {
	s := new(WSGTTGPDGService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGTTGPDGService() *WSGTTGPDGService {
	serv := new(WSGTTGPDGService)
	serv.apiClient = ss.apiClient
	return serv
}

// AddProfilesTtgpdg
//
// Use this API command to create TTG+PDG profile.
//
// Request Body:
//	 - body *profile.CreateTtgpdgProfile
func (s *WSGTTGPDGService) AddProfilesTtgpdg(ctx context.Context, body *profile.CreateTtgpdgProfile) (*common.CreateResult, error) {
}

// DeleteProfilesTtgpdg
//
// Use this API command to delete multiple TTG PDG profile.
//
// Request Body:
//	 - body *common.BulkDeleteRequest
func (s *WSGTTGPDGService) DeleteProfilesTtgpdg(ctx context.Context, body *common.BulkDeleteRequest) (*common.EmptyResult, error) {
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

// DeleteProfilesTtgpdgById
//
// Use this API command to delete TTG PDG profile.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGTTGPDGService) DeleteProfilesTtgpdgById(ctx context.Context, pId string) error {
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

// PartialUpdateProfilesTtgpdgById
//
// Use this API command to modify the basic information of TTG+PDG profile.
//
// Request Body:
//	 - body *profile.TtgpdgProfileConfiguration
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGTTGPDGService) PartialUpdateProfilesTtgpdgById(ctx context.Context, body *profile.TtgpdgProfileConfiguration, pId string) (*common.EmptyResult, error) {
}
