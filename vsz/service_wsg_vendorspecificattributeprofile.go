package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/vendorspecificattributeprofile"
)

type WSGVendorSpecificAttributeProfileService struct {
	apiClient *APIClient
}

func NewWSGVendorSpecificAttributeProfileService(c *APIClient) *WSGVendorSpecificAttributeProfileService {
	s := new(WSGVendorSpecificAttributeProfileService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGVendorSpecificAttributeProfileService() *WSGVendorSpecificAttributeProfileService {
	serv := new(WSGVendorSpecificAttributeProfileService)
	serv.apiClient = ss.apiClient
	return serv
}

// AddRkszonesVendorSpecificAttributeProfilesByZoneId
//
// Create a vendor specific attribute profile.
//
// Request Body:
//	 - body *vendorspecificattributeprofile.Persist
//
// Path Parameters:
// - pZoneId string
//		- required
func (s *WSGVendorSpecificAttributeProfileService) AddRkszonesVendorSpecificAttributeProfilesByZoneId(ctx context.Context, body *vendorspecificattributeprofile.Persist, pZoneId string) (vendorspecificattributeprofile.CreateResult, error) {
}

// DeleteRkszonesVendorSpecificAttributeProfilesById
//
// Use this API command to delete a vendor specific attribute profile by ID.
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGVendorSpecificAttributeProfileService) DeleteRkszonesVendorSpecificAttributeProfilesById(ctx context.Context, pId string, pZoneId string) (*vendorspecificattributeprofile.EmptyResult, error) {
}

// DeleteRkszonesVendorSpecificAttributeProfilesByZoneId
//
// Use this API command to delete a list of vendor specific attribute profile.
//
// Request Body:
//	 - body *vendorspecificattributeprofile.DeleteBulk
//
// Path Parameters:
// - pZoneId string
//		- required
func (s *WSGVendorSpecificAttributeProfileService) DeleteRkszonesVendorSpecificAttributeProfilesByZoneId(ctx context.Context, body *vendorspecificattributeprofile.DeleteBulk, pZoneId string) (*vendorspecificattributeprofile.EmptyResult, error) {
}

// FindRkszonesVendorSpecificAttributeProfilesById
//
// Get a vendor specific attribute profile by ID.
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGVendorSpecificAttributeProfileService) FindRkszonesVendorSpecificAttributeProfilesById(ctx context.Context, pId string, pZoneId string) (*vendorspecificattributeprofile.Get, error) {
}

// FindRkszonesVendorSpecificAttributeProfilesByQueryCriteria
//
// Use this API command to retrieve a list of vendor specific attribute profile by query criteria.
//
// Request Body:
//	 - body *common.QueryCriteriaSuperSet
func (s *WSGVendorSpecificAttributeProfileService) FindRkszonesVendorSpecificAttributeProfilesByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) (*vendorspecificattributeprofile.QueryCriteriaResult, error) {
}

// FindRkszonesVendorSpecificAttributeProfilesByZoneId
//
// Get a ID list of vendor specific attribute profile in this Zone.
//
// Path Parameters:
// - pZoneId string
//		- required
func (s *WSGVendorSpecificAttributeProfileService) FindRkszonesVendorSpecificAttributeProfilesByZoneId(ctx context.Context, pZoneId string) (*vendorspecificattributeprofile.List, error) {
}

// UpdateRkszonesVendorSpecificAttributeProfilesById
//
// Use this API command to modify entire information of a vendor specific attribute profile.
//
// Request Body:
//	 - body *vendorspecificattributeprofile.Persist
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGVendorSpecificAttributeProfileService) UpdateRkszonesVendorSpecificAttributeProfilesById(ctx context.Context, body *vendorspecificattributeprofile.Persist, pId string, pZoneId string) (*vendorspecificattributeprofile.EmptyResult, error) {
}
