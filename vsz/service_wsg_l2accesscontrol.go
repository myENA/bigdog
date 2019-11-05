package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/portalservice"
)

type WSGL2AccessControlService struct {
	apiClient *APIClient
}

func NewWSGL2AccessControlService(c *APIClient) *WSGL2AccessControlService {
	s := new(WSGL2AccessControlService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGL2AccessControlService() *WSGL2AccessControlService {
	serv := new(WSGL2AccessControlService)
	serv.apiClient = ss.apiClient
	return serv
}

// AddRkszonesL2ACLByZoneId
//
// Create a new L2 Access Control.
//
// Request Body:
//	 - body *portalservice.CreateL2ACL
//
// Path Parameters:
// - pZoneId string
//		- required
func (s *WSGL2AccessControlService) AddRkszonesL2ACLByZoneId(ctx context.Context, body *portalservice.CreateL2ACL, pZoneId string) (*common.CreateResult, error) {
}

// DeleteRkszonesL2ACLById
//
// Delete an L2 Access Control.
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGL2AccessControlService) DeleteRkszonesL2ACLById(ctx context.Context, pId string, pZoneId string) (*common.EmptyResult, error) {
}

// FindRkszonesL2ACLById
//
// Retrieve an L2 Access Control.
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGL2AccessControlService) FindRkszonesL2ACLById(ctx context.Context, pId string, pZoneId string) (*portalservice.L2ACL, error) {
}

// FindRkszonesL2ACLByZoneId
//
// Retrieve a list of L2 Access Control.
//
// Path Parameters:
// - pZoneId string
//		- required
//
// Query Parameters:
// - qIndex string
// - qListSize string
func (s *WSGL2AccessControlService) FindRkszonesL2ACLByZoneId(ctx context.Context, pZoneId string, qIndex string, qListSize string) (*portalservice.PortalServiceList, error) {
}

// PartialUpdateRkszonesL2ACLById
//
// Modify a specific L2 Access Control basic.
//
// Request Body:
//	 - body *portalservice.ModifyL2ACL
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGL2AccessControlService) PartialUpdateRkszonesL2ACLById(ctx context.Context, body *portalservice.ModifyL2ACL, pId string, pZoneId string) (*common.EmptyResult, error) {
}
