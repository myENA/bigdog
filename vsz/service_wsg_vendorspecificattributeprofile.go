package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/vendorspecificattributeprofile"
)

type WSGVendorSpecificAttributeProfileService struct {
	client *Client
}

func NewWSGVendorSpecificAttributeProfileService(client *Client) *WSGVendorSpecificAttributeProfileService {
	s := new(WSGVendorSpecificAttributeProfileService)
	s.client = client
	return s
}

func (ss *WSGService) WSGVendorSpecificAttributeProfileService() *WSGVendorSpecificAttributeProfileService {
	serv := new(WSGVendorSpecificAttributeProfileService)
	serv.client = ss.client
	return serv
}

func (s *WSGVendorSpecificAttributeProfileService) FindRkszonesVendorSpecificAttributeProfilesById(ctx context.Context, pId string, pZoneId string) (*vendorspecificattributeprofile.Get, error) {
}

func (s *WSGVendorSpecificAttributeProfileService) FindRkszonesVendorSpecificAttributeProfilesByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) (*vendorspecificattributeprofile.QueryCriteriaResult, error) {
}

func (s *WSGVendorSpecificAttributeProfileService) FindRkszonesVendorSpecificAttributeProfilesByZoneId(ctx context.Context, pZoneId string) (*vendorspecificattributeprofile.List, error) {
}
