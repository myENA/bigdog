package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/switchm/aclconfig"
	"github.com/myENA/ruckus-client/vsz/types/switchm/common"
)

type SwitchMSwitchAccessControlListService struct {
	apiClient *APIClient
}

func NewSwitchMSwitchAccessControlListService(c *APIClient) *SwitchMSwitchAccessControlListService {
	s := new(SwitchMSwitchAccessControlListService)
	s.apiClient = c
	return s
}

func (ss *SwitchMService) SwitchMSwitchAccessControlListService() *SwitchMSwitchAccessControlListService {
	serv := new(SwitchMSwitchAccessControlListService)
	serv.apiClient = ss.apiClient
	return serv
}

// AddAccessControls
//
// Use this API command to Create the Access Control Config.
//
// Request Body:
//	 - body *aclconfig.CreateACLConfig
func (s *SwitchMSwitchAccessControlListService) AddAccessControls(ctx context.Context, body *aclconfig.CreateACLConfig) (*common.CreateResult, error) {
}

// DeleteAccessControls
//
// Use this API command to Delete the Access Control Config by Id list.
//
// Request Body:
//	 - body *common.BulkDeleteRequest
func (s *SwitchMSwitchAccessControlListService) DeleteAccessControls(ctx context.Context, body *common.BulkDeleteRequest) error {
}

// DeleteAccessControlsById
//
// Use this API command to Delete the Access Control Config.
//
// Path Parameters:
// - pId string
//		- required
func (s *SwitchMSwitchAccessControlListService) DeleteAccessControlsById(ctx context.Context, pId string) error {
}

// FindAccessControlsById
//
// Use this API command to Retrieve the Access Control Config.
//
// Path Parameters:
// - pId string
//		- required
func (s *SwitchMSwitchAccessControlListService) FindAccessControlsById(ctx context.Context, pId string) (*aclconfig.ACLConfig, error) {
}

// FindAccessControlsByQueryCriteria
//
// Use this API command to Retrieve the Access Control Config list.
//
// Request Body:
//	 - body *common.QueryCriteriaSuperSet
func (s *SwitchMSwitchAccessControlListService) FindAccessControlsByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) (*aclconfig.ACLConfigsQueryResult, error) {
}

// UpdateAccessControlsById
//
// Use this API command to Update the Access Control Config.
//
// Request Body:
//	 - body *aclconfig.UpdateACLConfig
//
// Path Parameters:
// - pId string
//		- required
func (s *SwitchMSwitchAccessControlListService) UpdateAccessControlsById(ctx context.Context, body *aclconfig.UpdateACLConfig, pId string) (*aclconfig.EmptyResult, error) {
}
