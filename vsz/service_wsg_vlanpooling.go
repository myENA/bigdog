package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/vlanpooling"
)

type WSGVlanPoolingService struct {
	apiClient *APIClient
}

func NewWSGVlanPoolingService(c *APIClient) *WSGVlanPoolingService {
	s := new(WSGVlanPoolingService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGVlanPoolingService() *WSGVlanPoolingService {
	serv := new(WSGVlanPoolingService)
	serv.apiClient = ss.apiClient
	return serv
}

// AddVlanpoolings
//
// Use this API command to create new VLAN pooling.
//
// Request Body:
//	 - body *vlanpooling.CreateVlanPooling
func (s *WSGVlanPoolingService) AddVlanpoolings(ctx context.Context, body *vlanpooling.CreateVlanPooling) (*common.CreateResult, error) {
}

// DeleteVlanpoolings
//
// Use this API command to bulk delete VLAN pooling.
//
// Request Body:
//	 - body *vlanpooling.DeleteBulkVlanPooling
func (s *WSGVlanPoolingService) DeleteVlanpoolings(ctx context.Context, body *vlanpooling.DeleteBulkVlanPooling) (*common.EmptyResult, error) {
}

// DeleteVlanpoolingsById
//
// Use this API command to delete VLAN pooling.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGVlanPoolingService) DeleteVlanpoolingsById(ctx context.Context, pId string) (*common.EmptyResult, error) {
}

// FindVlanpoolingsById
//
// Use this API command to retrieve VLAN pooling.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGVlanPoolingService) FindVlanpoolingsById(ctx context.Context, pId string) (*vlanpooling.VlanPooling, error) {
}

// FindVlanpoolingsByQueryCriteria
//
// Use this API command to retrieve a list of VLAN poolings.
//
// Request Body:
//	 - body *common.QueryCriteriaSuperSet
func (s *WSGVlanPoolingService) FindVlanpoolingsByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) (*vlanpooling.VlanPoolingList, error) {
}

// PartialUpdateVlanpoolingsById
//
// Use this API command to modify the basic information on VLAN pooling.
//
// Request Body:
//	 - body *vlanpooling.ModifyVlanPooling
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGVlanPoolingService) PartialUpdateVlanpoolingsById(ctx context.Context, body *vlanpooling.ModifyVlanPooling, pId string) (*common.EmptyResult, error) {
}
