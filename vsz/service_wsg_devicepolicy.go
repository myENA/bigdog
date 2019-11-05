package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/devicepolicy"
)

type WSGDevicePolicyService struct {
	apiClient *APIClient
}

func NewWSGDevicePolicyService(c *APIClient) *WSGDevicePolicyService {
	s := new(WSGDevicePolicyService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGDevicePolicyService() *WSGDevicePolicyService {
	serv := new(WSGDevicePolicyService)
	serv.apiClient = ss.apiClient
	return serv
}

// AddRkszonesDevicePolicyByZoneId
//
// Create a new Device Policy Porfile.
//
// Request Body:
//	 - body *devicepolicy.CreateDevicePolicy
//
// Path Parameters:
// - pZoneId string
//		- required
func (s *WSGDevicePolicyService) AddRkszonesDevicePolicyByZoneId(ctx context.Context, body *devicepolicy.CreateDevicePolicy, pZoneId string) (*common.CreateResult, error) {
}

// DeleteRkszonesDevicePolicyById
//
// Delete Device Policy Porfile.
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGDevicePolicyService) DeleteRkszonesDevicePolicyById(ctx context.Context, pId string, pZoneId string) (*common.EmptyResult, error) {
}

// FindRkszonesDevicePolicyById
//
// Retrieve a Device Policy Porfile.
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGDevicePolicyService) FindRkszonesDevicePolicyById(ctx context.Context, pId string, pZoneId string) (*devicepolicy.DevicePolicyPorfile, error) {
}

// FindRkszonesDevicePolicyByZoneId
//
// Retrieve a list of Device Policy Porfiles within a zone.
//
// Path Parameters:
// - pZoneId string
//		- required
//
// Query Parameters:
// - qIndex string
// - qListSize string
func (s *WSGDevicePolicyService) FindRkszonesDevicePolicyByZoneId(ctx context.Context, pZoneId string, qIndex string, qListSize string) (*devicepolicy.PorfileList, error) {
}

// PartialUpdateRkszonesDevicePolicyById
//
// Modify a specific Device Policy Porfile.
//
// Request Body:
//	 - body *devicepolicy.ModifyDevicePolicy
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGDevicePolicyService) PartialUpdateRkszonesDevicePolicyById(ctx context.Context, body *devicepolicy.ModifyDevicePolicy, pId string, pZoneId string) (*common.EmptyResult, error) {
}
