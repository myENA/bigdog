package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/vlanpooling"
)

type WSGVlanPoolingService struct {
	client *Client
}

func NewWSGVlanPoolingService(client *Client) *WSGVlanPoolingService {
	s := new(WSGVlanPoolingService)
	s.client = client
	return s
}

func (ss *WSGService) WSGVlanPoolingService() *WSGVlanPoolingService {
	serv := new(WSGVlanPoolingService)
	serv.client = ss.client
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
