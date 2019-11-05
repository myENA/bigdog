package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/switchm/aaaservers"
	"github.com/myENA/ruckus-client/vsz/types/switchm/common"
)

type SwitchMSwitchAAAServersService struct {
	apiClient *APIClient
}

func NewSwitchMSwitchAAAServersService(c *APIClient) *SwitchMSwitchAAAServersService {
	s := new(SwitchMSwitchAAAServersService)
	s.apiClient = c
	return s
}

func (ss *SwitchMService) SwitchMSwitchAAAServersService() *SwitchMSwitchAAAServersService {
	serv := new(SwitchMSwitchAAAServersService)
	serv.apiClient = ss.apiClient
	return serv
}

// AddAaaServersAdmin
//
// Use this API command to create a new AAA server.
//
// Request Body:
//	 - body *aaaservers.CreateAdminAAAServer
func (s *SwitchMSwitchAAAServersService) AddAaaServersAdmin(ctx context.Context, body *aaaservers.CreateAdminAAAServer) (*common.CreateResult, error) {
}

// DeleteAaaServersAdmin
//
// Use this API command to delete AAA Servers.
//
// Request Body:
//	 - body *common.BulkDeleteRequest
func (s *SwitchMSwitchAAAServersService) DeleteAaaServersAdmin(ctx context.Context, body *common.BulkDeleteRequest) error {
}

// DeleteAaaServersAdminById
//
// Use this API command to delete a AAA server.
//
// Path Parameters:
// - pId string
//		- required
func (s *SwitchMSwitchAAAServersService) DeleteAaaServersAdminById(ctx context.Context, pId string) (*aaaservers.EmptyResult, error) {
}

// FindAaaServersAdmin
//
// Use this API command to retrieve a list of AAA server.
func (s *SwitchMSwitchAAAServersService) FindAaaServersAdmin(ctx context.Context) (*aaaservers.AaaServersQueryResult, error) {
}

// FindAaaServersAdminById
//
// Use this API command to retrieve a AAA server.
//
// Path Parameters:
// - pId string
//		- required
func (s *SwitchMSwitchAAAServersService) FindAaaServersAdminById(ctx context.Context, pId string) (*aaaservers.AAAServer, error) {
}

// UpdateAaaServersAdminById
//
// Use this API command to modify the basic information on a AAA server by complete attributes.
//
// Request Body:
//	 - body *aaaservers.CreateAdminAAAServer
//
// Path Parameters:
// - pId string
//		- required
func (s *SwitchMSwitchAAAServersService) UpdateAaaServersAdminById(ctx context.Context, body *aaaservers.CreateAdminAAAServer, pId string) (*aaaservers.EmptyResult, error) {
}
